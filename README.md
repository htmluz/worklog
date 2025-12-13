#### fluxos
nova window trackeada: cria *Task* + *Window*, chama window.Start()

nova window filha: cria *Window*, adiciona a *Task* existente, chama window.Start()

navega pra outra window: window.Pause() na antiga, window.Resume() na nova(se trackeada)

fecha window: window.Close(), task.RemoveWindow(), se ultima -> task.Close()

Status: calcula window.Duration() do interval ativo, da pra add quantia de interval dps


#### TODO

WORKDIR=~/.worklog

- [x] Storage Json - data.json
- [ ] State Runtime - active.json - vou deixa p dps mo preguica
- [ ] CLI
- [ ] Fluxos

tenho que pensar no service la dps que fiz to confuso referente ao fluxo que vai ser dos nomes do tmux
pensar se vai ser ao abrir pq quando abrir sempre vai ser msm nome tlgd, se eu gerar um id random como diabos eu vou saber que é a mesma
muitas duvidas...
