package main

import (
	"flag"
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"jtracer"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)
import _ "net/http/pprof"

const (
	padding  = 2
	maxWidth = 80
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

type tickMsg time.Time

type model struct {
	startTime    time.Time
	outputFile   string
	scene        *jtracer.Scene
	progressChan chan float64
	percent      float64
	progress     progress.Model
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tickCmd(), tea.EnterAltScreen)

}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 4
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case tickMsg:
		m.percent = <-m.progressChan
		if m.percent >= 1.0 {
			m.percent = 1.0
			return m, tea.Quit
		}
		return m, tickCmd()

	default:
		return m, nil
	}
}

var (
	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
)

const template = `
Input File:  %v
Output File: %v
Resolution: %v x %v
`

const background = " △ ▢ ◯"

func (m model) View() string {
	title := lipgloss.NewStyle().Width(80).Align(lipgloss.Center).Render(m.scene.Description.Title)
	details := lipgloss.NewStyle().Width(60).Align(lipgloss.Left).
		Render(fmt.Sprintf(template, m.scene.InputFile, m.outputFile, m.scene.Camera.Vsize, m.scene.Camera.Hsize))
	ui := lipgloss.JoinVertical(lipgloss.Center, title, details)

	elapsedTime := time.Now().Sub(m.startTime).Round(1 * time.Second)
	pad := strings.Repeat(" ", padding)
	paddedProgress := pad + m.progress.ViewAs(m.percent) + pad + "\n"
	//paddedProgress += pad + fmt.Sprintf("%v", elapsedTime) + "\n"
	paddedProgress += lipgloss.NewStyle().Width(80).Align(lipgloss.Left).
		Render(fmt.Sprintf("\nElapsed: %v", elapsedTime))
	//paddedProgress += pad + helpStyle("Press any key to quit   ")

	ui = lipgloss.JoinVertical(lipgloss.Center, ui, paddedProgress)

	dialogBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#874BFD")).
		Padding(1, 0).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)

	dialog := lipgloss.Place(98, 20,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars(background),
		lipgloss.WithWhitespaceForeground(subtle),
	)
	return dialog

}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func main() {
	height := flag.Float64("height", -1, "Height of output image")
	width := flag.Float64("width", -1, "Height of output image")
	outputFile := flag.String("out", "out.png", "Filename of output image")

	flag.Parse()

	inputFileName := os.Args[len(os.Args)-1]

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	if *height != -1.0 {
		// TODO
	}

	if *width != -1.0 {
		// TODO
	}

	scene, err := jtracer.LoadSceneFile(inputFileName)
	if err != nil {
		panic(err)
	}

	go func() {
		canvas := scene.Camera.Render(jtracer.World{
			Objects: scene.Objects,
			Light:   scene.Light,
		})
		err = canvas.SavePNG(*outputFile)
	}()

	prog := progress.New(progress.WithDefaultGradient())

	_, err = tea.NewProgram(model{
		startTime:    time.Now(),
		outputFile:   *outputFile,
		scene:        scene,
		progress:     prog,
		progressChan: scene.Camera.Progress,
	}, tea.WithAltScreen()).Run()
	if err != nil {
		panic(err)
	}
}
