class Person {
    constructor(name?: string) {
        if (name == undefined) {
            console.log(`Person constructor without name`);
        } else {
            console.log(`Person constructor with name = ${name}`);
        }
    }

}

const personInstance = new Person("Aaron");
const unknownPerson = new Person;

