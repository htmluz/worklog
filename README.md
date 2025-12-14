#### fluxos
nova window trackeada: cria **Task** + **Window**, chama window.Start()

nova window filha: cria **Window**, adiciona a **Task** existente, chama window.Start()

navega pra outra window: window.Pause() na antiga, window.Resume() na nova(se trackeada)

fecha window: window.Close(), task.RemoveWindow(), se ultima -> task.Close()

Status: calcula window.Duration() do interval ativo, da pra add quantia de interval dps


#### TODO

WORKDIR=~/.worklog

- [x] Storage Json - data.json
- [ ] State Runtime - active.json - vou deixa p dps mo preguica
- [ ] CLI
- [ ] Fluxos


## My little pain in the ass with jira hours

i dont remember shit ive done and have to create my tasks and put the hours in there so im doing this to hook it up with my tmux with the faith it will save me.

This is all done following the best practices ever like DRY(DO repeat yourself) and [XGH](https://gohorse.com.br/extreme-go-horse-xgh.html)
