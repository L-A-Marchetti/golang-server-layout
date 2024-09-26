package handlers

import "config"

type Index struct {
	PageTitle string
	Header    Header
	Footer    Footer
}

type Header struct {
	Title string
	Links []Link
}

type Link struct {
	Href string
	Text string
}

type Footer struct {
	SiteName string
}

var (
	HeaderData = Header{
		Title: config.WEBSITE_TITLE,
		Links: []Link{
			{
				Href: "/",
				Text: "Home",
			},
			{
				Href: "/about",
				Text: "About",
			},
			{
				Href: "#contact",
				Text: "Contact",
			},
		},
	}

	FooterData = Footer{
		SiteName: config.WEBSITE_TITLE,
	}

	IndexData = Index{
		PageTitle: "Home | " + config.WEBSITE_TITLE,
		Header:    HeaderData,
		Footer:    FooterData,
	}
)
