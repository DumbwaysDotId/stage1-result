// Encapsulation
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

    set make(value) {
        if (value === "") {
            console.log("The make cannot be an empty string.");
            return;
        }
        this.#make = value;
    }

    get model() {
        return this.#model;
    }

    set model(value) {
        if (value === "") {
            console.log("The model cannot be an empty string.");
            return;
        }
        this.#model = value;
    }

    drive() {
        console.log(`Driving a ${this.#make} ${this.#model}.`);
    }
}

let car = new Car("Toyota", "Camry");
console.log(car.make); // Output: Toyota
car.make = "Honda";
console.log(car.make); // Output: Honda

car.drive(); // Output: Driving a Honda Camry.


