package congo

//ActionMap -
var ActionMap TActionMap

//TActionMap -
type TActionMap struct {
	byName    map[string]IAction
	byTrigger map[string]IAction
	state     string
}

/////////////////////////////

//IAction -
type IAction interface {
	GetName() string
	GetTrigger() string
	GetDescr() string
	Do(ev IEvent) bool
}

//TAbstractAction -
type TAbstractAction struct {
	name        string
	trigger     string
	description string
	//handler     func(ev IEvent) bool
}

//GetName -
func (act *TAbstractAction) GetName() string {
	return act.name
}

//GetTrigger -
func (act *TAbstractAction) GetTrigger() string {
	return act.trigger
}

//GetDescr -
func (act *TAbstractAction) GetDescr() string {
	return act.description
}

//Do -
func (act *TAbstractAction) Do(ev IEvent) bool {
	panic("Abstract call Func")
	return false
}

/////////////////////////////////////////////////

//TAction -
type TAction struct {
	TAbstractAction
	handler func(ev IEvent) bool
}

//Do -
func (act *TAction) Do(ev IEvent) bool {
	return act.handler(ev)
}

//NewAction -
func NewAction(name, trigger, description string, handler func(ev IEvent) bool) IAction {
	action := &TAction{}
	action.name = name
	action.trigger = trigger
	action.description = description
	action.handler = handler
	ActionMap.byName[action.name] = action
	return action
}

//////////////////////////////////////////////////

//TKeyboardAction -
type TKeyboardAction struct {
	TAbstractAction
	handler func(ev *KeyboardEvent) bool
}

//Do -
func (act *TKeyboardAction) Do(ev IEvent) bool {
	kbd, ok := ev.(*KeyboardEvent)
	if !ok {
		//TODO !!! error msg log
		return false
	}
	return act.handler(kbd)
}

//NewKeyboardAction -
func NewKeyboardAction(name, trigger, description string, handler func(ev *KeyboardEvent) bool) IAction {
	action := &TKeyboardAction{}
	action.name = name
	action.trigger = trigger
	action.description = description
	action.handler = handler
	ActionMap.byName[action.name] = action
	return action
}

///////////////////////////////////////////////////

//TResizeAction -
type TResizeAction struct {
	TAbstractAction
	handler func(ev *ResizeEvent) bool
}

//Do -
func (act *TResizeAction) Do(ev IEvent) bool {
	cons, ok := ev.(*ResizeEvent)
	if !ok {
		//TODO !!! error msg log
		return false
	}
	return act.handler(cons)
}

//NewResizeAction -
func NewResizeAction(name, trigger, description string, handler func(ev *ResizeEvent) bool) IAction {
	action := &TResizeAction{}
	action.name = name
	action.trigger = trigger
	action.description = description
	action.handler = handler
	ActionMap.byName[action.name] = action
	return action
}

///////////////////////////////////////////////////

//TMouseAction -
type TMouseAction struct {
	TAbstractAction
	handler func(ev *MouseEvent) bool
}

//Do -
func (act *TMouseAction) Do(ev IEvent) bool {
	cons, ok := ev.(*MouseEvent)
	if !ok {
		//TODO !!! error msg log
		return false
	}
	return act.handler(cons)
}

//NewMouseAction -
func NewMouseAction(name, trigger, description string, handler func(ev *MouseEvent) bool) IAction {
	action := &TMouseAction{}
	action.name = name
	action.trigger = trigger
	action.description = description
	action.handler = handler
	ActionMap.byName[action.name] = action
	return action
}

///////////////////////////////////////////////////

func initActionMap() {
	ActionMap = TActionMap{}
	ActionMap.byName = map[string]IAction{}
}

//GetNames - 
func (am *TActionMap) GetNames() []string {
	names := make([]string, 0, len(am.byName))
	for k, act := range am.byName {
		names = append(names, act.GetTrigger()+" - "+k)
	}
	return names
}

/*func (am *tActionMap) Names() []string {
  names := make([]string, 0, len(am.byName))
  for k, act := range am.byName {
    names = append(names, act.EventKey()+" - "+k)
  }
  return names
}*/

//Apply -
func (am *TActionMap) Apply() {
	am.byTrigger = map[string]IAction{}
	for _, act := range am.byName {
		if _, ok := am.byTrigger[act.GetTrigger()]; ok {
			//!!! TODO - log warning
		}
		am.byTrigger[act.GetTrigger()] = act
	}
}

//SetState -
func (am *TActionMap) SetState(state string) {
	am.state = state
	if state != "" {
		am.state += "/"
	}
}

//GetState - 
func (am *TActionMap) GetState() string {
	return am.state
}

//HandleEvent -
func HandleEvent(ev IEvent) {
	isHandled := false
	action, ok := ActionMap.byTrigger[ActionMap.GetState()+ev.GetTrigger()]
	if ok {
		isHandled = action.Do(ev)
	}
	if isHandled {
		return
	}
	action, ok = ActionMap.byTrigger[ev.GetTrigger()]
	if ok {
		action.Do(ev)
	}

}
