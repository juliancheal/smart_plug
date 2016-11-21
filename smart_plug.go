package smart_plug

type SmartPlug interface {
    TurnOn() error
    TurnOff() error
}
