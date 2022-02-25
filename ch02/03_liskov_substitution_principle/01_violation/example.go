package lsp_violation

func Go(vehicle actions) {
	if sled, ok := vehicle.(*Sled); ok {
		sled.pushStart()
	} else {
		vehicle.startEngine()
	}

	vehicle.drive()
}

type actions interface {
	drive()
	startEngine()
}

type Vehicle struct {
}

func (v Vehicle) drive() {
	// TODO: implement
}

func (v Vehicle) startEngine() {
	// TODO: implement
}

func (v Vehicle) stopEngine() {
	// TODO: implement
}

type Car struct {
	Vehicle
}

type Sled struct {
	Vehicle
}

func (s Sled) startEngine() {
	// override so that is does nothing
}

func (s Sled) stopEngine() {
	// override so that is does nothing
}

func (s Sled) pushStart() {
	// TODO: implement
}







type TheAx interface {
	cut()
}

type axOne struct {

}

func (a axOne) cut()  {

}

type PowerAx struct {
	axOne
}

func test(a TheAx)  {

}

type PoweAxII struct {
	PowerAx
}

func (i PoweAxII) cut()  {
	
}

func caller() {
	pa := PoweAxII{}
	test(pa)
}