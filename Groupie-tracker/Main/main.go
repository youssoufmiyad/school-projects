package main

import (
	"fmt"
	fonction "groupie/Fonction"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Définition du theme
type myTheme struct{}

func (myTheme) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch c {
	case theme.ColorNameBackground:
		return color.NRGBA{R: 0x3b, G: 0x11, B: 0x77, A: 0x90}
	case theme.ColorNameMenuBackground:
		return color.NRGBA{R: 0x3b, G: 0x11, B: 0x77, A: 0x90}
	case theme.ColorNameButton:
		return color.NRGBA{R: 0xce, G: 0x9a, B: 0x24, A: 0xfa}
	case theme.ColorNameDisabledButton:
		return color.NRGBA{R: 0xa5, G: 0x72, B: 0x34, A: 0xbf}
	case theme.ColorNameDisabled:
		return color.NRGBA{R: 0xce, G: 0x91, B: 0x24, A: 0xfa}
	case theme.ColorNameError:
		return color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff}
	case theme.ColorNameFocus:
		return color.NRGBA{R: 0xce, G: 0x9a, B: 0x24, A: 0xfa}
	case theme.ColorNameForeground:
		return color.NRGBA{R: 0xe1, G: 0xba, B: 0x66, A: 0xf9}
	case theme.ColorNameHover:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xf}
	case theme.ColorNameInputBackground:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x19}
	case theme.ColorNamePlaceHolder:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x19}
	case theme.ColorNamePressed:
		return color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x66}
	case theme.ColorNamePrimary:
		return color.NRGBA{R: 0xce, G: 0x9a, B: 0x24, A: 0xfa}
	case theme.ColorNameScrollBar:
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x99}
	case theme.ColorNameShadow:
		return color.NRGBA{R: 0x0, G: 0x0, B: 0x0, A: 0x66}
	default:
		return theme.DefaultTheme().Color(c, v)
	}
}

func (myTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return theme.DefaultTheme().Font(s)
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return theme.DefaultTheme().Font(s)
}

func (myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (myTheme) Size(s fyne.ThemeSizeName) float32 {
	switch s {
	case theme.SizeNameCaptionText:
		return 11
	case theme.SizeNameInlineIcon:
		return 20
	case theme.SizeNamePadding:
		return 4
	case theme.SizeNameScrollBar:
		return 16
	case theme.SizeNameScrollBarSmall:
		return 3
	case theme.SizeNameSeparatorThickness:
		return 1
	case theme.SizeNameText:
		return 14
	case theme.SizeNameInputBorder:
		return 2
	default:
		return theme.DefaultTheme().Size(s)
	}
}

var id_nbr int
var Nom_tab []string
var Dates_tab []string
var Locat_tab []string
var Relation_lieu_tab []string
var Bigger_tab []string
var Total_tab [][]string
var options []string
var Dt string
var Lt string

func main() {
	// récolte API
	fonction.GetArtists()
	fonction.GetDates()
	fonction.GetLocations()
	fonction.GetRelations()

	// initialisation de l'appli
	myApp := app.New()
	myApp.Settings().SetTheme(&myTheme{})

	// page principal
	icon, _ := fyne.LoadResourceFromPath("icon.png")
	menu := myApp.NewWindow("Menu")
	menu.Resize(fyne.NewSize(400, 500))
	menu.CenterOnScreen()
	menu.SetIcon(icon)

	// page d'info
	info := myApp.NewWindow("info")
	info.Resize(fyne.NewSize(700, 520))
	info.CenterOnScreen()
	info.SetIcon(icon)

	// chargement image
	id_nbr = 1
	r, _ := fyne.LoadResourceFromURLString(fonction.Artists[id_nbr].Image)
	img := canvas.NewImageFromResource(r)
	img.Resize(fyne.NewSize(300, 300))
	img.Move(fyne.NewPos(400, 50))
	info.SetIcon(img.Resource)

	// création list des noms de groupes
	for i := 0; i < 52; i++ {
		Nom_tab = append(Nom_tab, fonction.Artists[i].Name)
	}

	// bouton de sortie de info
	exit := widget.NewButton("exit", func() {
		info.Hide()
		menu.Show()
	})
	exit.Resize(fyne.NewSize(450, 70))
	exit.Move(fyne.NewPos(250, 450))

	// barres de sélection contenant tout les lieux où il y a eu concert
	for i := range fonction.Relations.Index[id_nbr].Dates_Locations {
		Relation_lieu_tab = append(Relation_lieu_tab, i)
	}
	lieu := locations()

	// barres de sélection contenant toutes les infos complémentaires des groupes
	options = append(options, "Membres", "Début de carrière", "Premier album")
	description := infos()

	// liste contenant les dates de concerts
	date := widget.NewList(
		func() int {
			return len(fonction.Relations.Index[id_nbr].Dates_Locations[lieu.Selected])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("dates")
		}, func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(fonction.Relations.Index[id_nbr].Dates_Locations[lieu.Selected][lii])
		})
	date.Resize(fyne.NewSize(250, 200))
	date.Move(fyne.NewPos(0, 50))

	// affichage des noms de la list sur une colonne
	noms := noms()

	// cont_* = container contenant tout le contenu de la page "*"
	cont_info := container.NewWithoutLayout(lieu, date, description, img, exit)

	//input renvoyant vers la page info
	input := fonction.NewCompletionEntry([]string{})
	// on implémente l'auto complétion
	input.OnChanged = func(s string) {
		if len(s) < 3 {
			input.HideCompletion()
			return
		}
		var tab []string
		entrée := make([]string, len(input.Text))
		for i, r := range input.Text {
			entrée[i] = string(r)
		}
		for x := 0; x < len(fonction.Artists); x++ {
			// on va parcourir le tableau  et si on tombe sur nom correspondant on l'affiche
			comp_artist := make([]string, len(fonction.Artists[x].Name))
			for i, r := range fonction.Artists[x].Name {
				comp_artist[i] = string(r)
			}
			test_comp_artist := make([]string, len(input.Text))
			for i := 0; i < len(input.Text); i++ {
				//pour eviter les erreur et l'arret du programme ainsi les mots dont la taille est inferieur a celle du mot demandé
				if len(comp_artist) >= len(input.Text) {
					test_comp_artist[i] = comp_artist[i]
				}
			}
			fonction.WordToMin(entrée)
			fonction.WordToMin(test_comp_artist)
			text_entry := strings.Join(entrée, "")
			mot := strings.Join(test_comp_artist, "")
			fmt.Println(mot)
			if text_entry == mot {
				//ajoute le nom dans le tableau tab qui contiendra que les noms correspondants au nom
				tab = append(tab, fonction.Artists[x].Name)
			}
		}
		// Affichage de la complétion
		fmt.Println(tab)
		input.SetOptions(tab)
		input.ShowCompletion()
	}
	input.SetPlaceHolder("Entrer nom de l'artiste, lieu, membre...")
	input.Resize(fyne.NewSize(400, 80))
	input.Move(fyne.NewPos(0, 0))

	// recherche par nom de groupe/membre
	search := widget.NewButton("rechercher", func() {
		// print("\ntxt = ", input.Text, "\nid = ", id_nbr)
		for t := 0; t < len(fonction.Artists); t++ {
			if fonction.Artists[t].Name == input.Text {
				id_nbr = fonction.Artists[t].Id
				if id_nbr > 0 {
					id_nbr = id_nbr - 1
				}
				t = len(fonction.Artists)
				r, _ = fyne.LoadResourceFromURLString(fonction.Artists[id_nbr].Image)
				img = canvas.NewImageFromResource(r)
				img.Resize(fyne.NewSize(450, 450))
				img.Move(fyne.NewPos(250, 0))
				info.SetIcon(img.Resource)
				Relation_lieu_tab = nil
				for i := range fonction.Relations.Index[id_nbr].Dates_Locations {
					Relation_lieu_tab = append(Relation_lieu_tab, i)
				}
				lieu = locations()

				description = infos()

				lieu.OnChanged = func(s string) {
					date = widget.NewList(
						func() int {
							return len(fonction.Relations.Index[id_nbr].Dates_Locations[s])
						},
						func() fyne.CanvasObject {
							return widget.NewLabel("dates")
						}, func(lii widget.ListItemID, co fyne.CanvasObject) {
							co.(*widget.Label).SetText(fonction.Relations.Index[id_nbr].Dates_Locations[s][lii])
						})
					date.Resize(fyne.NewSize(250, 200))
					date.Move(fyne.NewPos(0, 50))
					cont_info = container.NewWithoutLayout(lieu, date, description, img, exit)
					info.SetContent(cont_info)
				}

				description.OnChanged = func(s string) {
					if s == "Membres" {
						desc := widget.NewList(
							func() int {
								return len(fonction.Artists[id_nbr].Members)
							},
							func() fyne.CanvasObject {
								return widget.NewLabel("membres")
							},
							func(lii widget.ListItemID, co fyne.CanvasObject) {
								co.(*widget.Label).SetText(fonction.Artists[id_nbr].Members[lii])
							})
						desc.Resize(fyne.NewSize(250, 200))
						desc.Move(fyne.NewPos(0, 300))
						cont_info = container.NewWithoutLayout(lieu, date, description, desc, img, exit)
						info.SetContent(cont_info)
					} else if s == "Début de carrière" {
						desc := widget.NewLabel(fmt.Sprint(fonction.Artists[id_nbr].CreationDate))
						desc.Resize(fyne.NewSize(250, 200))
						desc.Move(fyne.NewPos(0, 300))
						cont_info = container.NewWithoutLayout(lieu, date, description, desc, img, exit)
						info.SetContent(cont_info)
					} else if s == "Premier album" {
						desc := widget.NewLabel(fonction.Artists[id_nbr].FirstAlbum)
						desc.Resize(fyne.NewSize(250, 200))
						desc.Move(fyne.NewPos(0, 300))
						cont_info = container.NewWithoutLayout(lieu, date, description, desc, img, exit)
						info.SetContent(cont_info)
					}
				}

				cont_info = container.NewWithoutLayout(lieu, date, description, img, exit)
				menu.Hide()
				info.SetContent(cont_info)
				info.Show()
			} else {
				print("inconnu")
			}
		}
	})
	search.Resize(fyne.NewSize(400, 80))
	search.Move(fyne.NewPos(0, 80))

	cont_menu := container.NewWithoutLayout(noms, input, search)

	// interactivité des noms (ils agissent comme des boutons)
	noms.OnSelected = func(id widget.ListItemID) {
		id_nbr = id
		r, _ = fyne.LoadResourceFromURLString(fonction.Artists[id_nbr].Image)
		img = canvas.NewImageFromResource(r)
		img.Resize(fyne.NewSize(450, 450))
		img.Move(fyne.NewPos(250, 0))
		info.SetIcon(img.Resource)
		Relation_lieu_tab = nil
		for i := range fonction.Relations.Index[id_nbr].Dates_Locations {
			Relation_lieu_tab = append(Relation_lieu_tab, i)
		}
		lieu = locations()

		description := infos()

		lieu.OnChanged = func(s string) {
			date = widget.NewList(
				func() int {
					return len(fonction.Relations.Index[id_nbr].Dates_Locations[s])
				},
				func() fyne.CanvasObject {
					return widget.NewLabel("dates")
				}, func(lii widget.ListItemID, co fyne.CanvasObject) {
					co.(*widget.Label).SetText(fonction.Relations.Index[id_nbr].Dates_Locations[s][lii])
				})
			date.Resize(fyne.NewSize(250, 200))
			date.Move(fyne.NewPos(0, 50))
			cont_info = container.NewWithoutLayout(lieu, date, description, img, exit)
			info.SetContent(cont_info)
		}

		description.OnChanged = func(s string) {
			if s == "Membres" {
				desc := widget.NewList(
					func() int {
						return len(fonction.Artists[id_nbr].Members)
					},
					func() fyne.CanvasObject {
						return widget.NewLabel("membres")
					},
					func(lii widget.ListItemID, co fyne.CanvasObject) {
						co.(*widget.Label).SetText(fonction.Artists[id_nbr].Members[lii])
					})
				desc.Resize(fyne.NewSize(250, 300))
				desc.Move(fyne.NewPos(0, 300))
				cont_info = container.NewWithoutLayout(lieu, date, description, desc, img, exit)
				info.SetContent(cont_info)
			} else if s == "Début de carrière" {
				desc := widget.NewLabel(fmt.Sprint(fonction.Artists[id_nbr].CreationDate))
				desc.Resize(fyne.NewSize(250, 200))
				desc.Move(fyne.NewPos(0, 300))
				cont_info = container.NewWithoutLayout(lieu, date, description, desc, img, exit)
				info.SetContent(cont_info)
			} else if s == "Premier album" {
				desc := widget.NewLabel(fonction.Artists[id_nbr].FirstAlbum)
				desc.Resize(fyne.NewSize(250, 200))
				desc.Move(fyne.NewPos(0, 300))
				cont_info = container.NewWithoutLayout(lieu, date, description, desc, img, exit)
				info.SetContent(cont_info)
			}
		}

		cont_info = container.NewWithoutLayout(lieu, date, description, img, exit)
		menu.Hide()
		info.SetContent(cont_info)
		info.Show()
	}

	lieu.OnChanged = func(s string) {
		date = widget.NewList(
			func() int {
				return len(fonction.Relations.Index[id_nbr].Dates_Locations[s])
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("dates")
			}, func(lii widget.ListItemID, co fyne.CanvasObject) {
				co.(*widget.Label).SetText(fonction.Relations.Index[id_nbr].Dates_Locations[s][lii])
			})
		cont_info = container.NewWithoutLayout(lieu, date, img, exit)
		info.SetContent(cont_info)
	}

	menu.SetContent(cont_menu)
	menu.ShowAndRun()
}

func locations() *widget.Select {
	lieu := widget.NewSelect(Relation_lieu_tab, func(s string) {})
	lieu.Selected = "locations"
	lieu.Resize(fyne.NewSize(250, 50))
	lieu.Move(fyne.NewPos(0, 0))
	return lieu
}

func infos() *widget.Select {
	description := widget.NewSelect(options, func(s string) {})
	description.Selected = "infos"
	description.Resize(fyne.NewSize(250, 50))
	description.Move(fyne.NewPos(0, 250))
	return description
}

func noms() *widget.List {
	noms := widget.NewList(
		func() int {
			return 52
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("noms")
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(Nom_tab[lii])
		})
	noms.Resize(fyne.NewSize(400, 345))
	noms.Move(fyne.NewPos(0, 160))
	return noms
}
