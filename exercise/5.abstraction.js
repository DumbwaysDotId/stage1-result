// Abstraction:
class Car {
    #make = ""
    #model = ""

    constructor(make, model) {
        this.#make = make;
        this.#model = model;
    }

    get make() {
        return this.#make;
    }

    get model() {
        return this.#model;
    }
}

let myCar = new Car("Toyota", "Camry");
console.log(myCar.make); // Output: Toyota
console.log(myCar.model); // Output: Camry