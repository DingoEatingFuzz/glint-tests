package main

import (
  "context"

  "github.com/mitchellh/go-glint"
  cc "github.com/dingoeatingfuzz/glint-tests/components"
)

func main() {
  d := glint.New()

  d.Append(cc.WatchEvent(true, glint.Text("wow, we did it!"), []glint.Component{
    glint.Style(glint.Text("Red fish"), glint.Color("red")),
    glint.Style(glint.Text("Blue fish"), glint.Color("blue")),
  }))

  d.Append(cc.WatchEvent(false, glint.Text("wow, we did it!"), []glint.Component{
    glint.Style(glint.Text("Red fish"), glint.Color("red")),
    glint.Style(glint.Text("Blue fish"), glint.Color("blue")),
  }))

  d.Append(cc.LineSpacing())

  d.Append(glint.Layout(
    cc.Text("This is your standard text, which can be"),
    cc.Text(" bold").Bold(),
    cc.Text(" or "),
    cc.Text("underlined").Underline(),
  ).Row())

  d.Append(glint.Layout(
    cc.Subtle("This is your subtle text, which can be"),
    cc.Subtle(" bold").Bold(),
    cc.Subtle(" or "),
    cc.Subtle("underlined").Underline(),
  ).Row())

  d.Append(glint.Layout(
    cc.IconSuccess(),
    cc.Success("This is your success text, which can be"),
    cc.Success(" bold").Bold(),
    cc.Success(" or "),
    cc.Success("underlined").Underline(),
  ).Row())

  d.Append(glint.Layout(
    cc.IconHealthy(),
    cc.Info("This is your info text, which can be"),
    cc.Info(" bold").Bold(),
    cc.Info(" or "),
    cc.Info("underlined").Underline(),
  ).Row())

  d.Append(glint.Layout(
    cc.IconWarning(),
    cc.Warning("This is your warning text, which can be"),
    cc.Warning(" bold").Bold(),
    cc.Warning(" or "),
    cc.Warning("underlined").Underline(),
  ).Row())

  d.Append(glint.Layout(
    cc.IconFailure(),
    cc.Error("This is your error text, which can be"),
    cc.Error(" bold").Bold(),
    cc.Error(" or "),
    cc.Error("underlined").Underline(),
  ).Row())

  d.Append(cc.LineSpacing())

  d.Append(cc.LargeEvent(
    glint.Layout(
      cc.IconSuccess(),
      cc.Success("Healthy "),
      cc.Text("driver "),
      cc.Text("Docker").Bold(),
      cc.Text(" on client "),
      cc.Text("982ed12a").Bold(),
      cc.Subtle(" (ip-10.392.4.2)"),
    ).Row(),
    []glint.Component{
      glint.Layout(cc.Text("Could not connect to socket")).Row(),
      glint.Layout(cc.Subtle("To see more details, run: "), cc.Subtle("nomad node status 982ed12a").Bold()).Row(),
    },
  ))

  d.Append(cc.LineSpacing())

  d.Append(cc.SmallEvent(
    glint.Layout(cc.IconRunning(), cc.Success("Eligible   ")).Row(),
    cc.Text("client "),
    cc.Text("982ed12a").Bold(),
    cc.Subtle(" (ip-10.392.4.2)")),
  )

  d.Render(context.Background())
}
