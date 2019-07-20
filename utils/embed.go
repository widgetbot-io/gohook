package utils

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

type Embed struct {
	discordgo.MessageEmbed
}

func NewEmbed() *Embed {
	return &Embed{}
}

func (e *Embed) SetTitle(title string) *Embed {
	if len(title) > 256 {
		panic("Title supplied to EmbedBuilder is over 256 characters")
	}

	e.Title = title
	return e
}

func (e *Embed) SetDescription(description string) *Embed {
	if len(description) > 2048 {
		panic("Description supplied to EmbedBuilder is over 2048 characters")
	}

	e.Description = description
	return e
}

func (e *Embed) AddField(title string, content string, inline bool) *Embed {
	if len(title) > 1024 || len(content) > 1024 {
		panic("Title or content passed to EmbedBuilder is over 1024 characters.")
	}

	e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
		Name:   title,
		Value:  content,
		Inline: inline,
	})
	return e
}

func (e *Embed) SetURL(url string) *Embed {
	e.URL = url
	return e
}

func (e *Embed) SetTimestamp() *Embed {
	e.Timestamp = time.Now().Format(time.RFC3339)
	return e
}

func (e *Embed) SetColour(colour int) *Embed {
	e.Color = colour
	return e
}

func (e *Embed) SetFooter(title string, icon string) *Embed {
	e.Footer = &discordgo.MessageEmbedFooter{
		Text:    title,
		IconURL: icon,
	}
	return e
}

func (e *Embed) SetImage(url string) *Embed {
	e.Image = &discordgo.MessageEmbedImage{
		URL: url,
	}
	return e
}

func (e *Embed) SetThumbnail(url string) *Embed {
	e.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL: url,
	}
	return e
}

func (e *Embed) SetAuthor(name string, icon string) *Embed {
	e.Author = &discordgo.MessageEmbedAuthor{
		Name:    name,
		IconURL: icon,
		URL:     icon,
	}
	return e
}
