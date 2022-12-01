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
	termWidth    int
	termHeight   int
	startTime    time.Time
	endTime      *time.Time
	duration     time.Duration
	outputFile   string
	scene        *jtracer.Scene
	progressChan chan float64
	percent      float64
	progress     progress.Model
}

func (m *model) Init() tea.Cmd {
	go func() {
		for {
			m.percent = <-m.progressChan
		}
	}()

	return tea.Batch(tickCmd(), tea.EnterAltScreen)

}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
		m.progress.Width = msg.Width - 25
		//if m.progress.Width > maxWidth {
		//	m.progress.Width = maxWidth
		//}
		return m, nil

	case tickMsg:
		if m.percent >= 1.0 && m.endTime == nil {
			t := time.Now()
			m.endTime = &t
			m.percent = 1.0
			m.duration = time.Now().Sub(m.startTime).Round(1 * time.Second)
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
Resolution:  %v x %v`

const doneTemplate = `
File: %v
Render duration: %v
`

const background = " △ ▢ ◯"

func (m *model) doneDialog() string {
	ui := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().
			Italic(true).
			Align(lipgloss.Center).
			Foreground(lipgloss.Color("#FFF7DB")).
			//Background(lipgloss.Color("#F25D94")).
			Render("Render Complete!"),
		lipgloss.NewStyle().Width(60).Align(lipgloss.Left).
			Render(fmt.Sprintf(doneTemplate, m.outputFile, m.duration)),
		helpStyle("Press any key to exit"),
	)

	dialogBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Width(m.termWidth-20).
		BorderForeground(lipgloss.Color("46")).
		Padding(1, 2).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)

	return lipgloss.Place(
		m.termWidth,
		m.termHeight,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars(background),
		lipgloss.WithWhitespaceForeground(subtle),
	)
}

func (m *model) progressDialog() string {
	ui := lipgloss.JoinVertical(
		lipgloss.Center,
		lipgloss.NewStyle().
			Width(m.termWidth).
			Align(lipgloss.Left).
			Bold(true).
			Render(m.scene.Description.Title),
		lipgloss.NewStyle().
			Width(m.termWidth).
			Align(lipgloss.Left).
			Render(fmt.Sprintf(template, m.scene.InputFile, m.outputFile, m.scene.Camera.Hsize, m.scene.Camera.Vsize)),
	)

	paddedProgress := m.progress.ViewAs(m.percent) + "\n"
	paddedProgress += lipgloss.NewStyle().Width(m.termWidth).Align(lipgloss.Left).
		Render(fmt.Sprintf("\nElapsed: %v", time.Now().Sub(m.startTime).Round(1*time.Second)))

	ui = lipgloss.JoinVertical(lipgloss.Center, ui, paddedProgress)

	dialogBoxStyle := lipgloss.NewStyle().
		Width(m.termWidth-20).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#874BFD")).
		Padding(1, 2).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)

	dialog := lipgloss.Place(m.termWidth, m.termHeight,
		lipgloss.Center, lipgloss.Center,
		dialogBoxStyle.Render(ui),
		lipgloss.WithWhitespaceChars(background),
		lipgloss.WithWhitespaceForeground(subtle),
	)
	return dialog
}

func (m *model) View() string {
	if m.percent >= 1.0 {
		return m.doneDialog()
	}

	return m.progressDialog()
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

	_, err = tea.NewProgram(&model{
		startTime:    time.Now(),
		outputFile:   *outputFile,
		scene:        scene,
		progress:     progress.New(progress.WithDefaultGradient()),
		progressChan: scene.Camera.Progress,
	}, tea.WithAltScreen()).Run()
	if err != nil {
		panic(err)
	}

}
