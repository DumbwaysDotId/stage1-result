// Step : 
// - making class
// - making object
// - called it so it appear in browser using loop and innerHTML
// - refactor the code so it implement encapsulation, use getter and setter(the key point is : encapsulation make app more secure by implementing restriction)
// - refactor the code so it implements inheritance AuthorTestimonial
// - refactor the code so it implementes polymorphism by making class CompanyTestimonial and override get author(the key point is polymorphism implement and overriding)
// - refactor code so it implements abstraction by making get testimonialHTML and get author throw error when children didn't use it

// class Testimonial {
//     #quote = ""
//     #image = ""

//     constructor(quote, image) {
//         this.#quote = quote;
//         this.#image = image;
//     }

//     get quote() {
//         return this.#quote;
//     }

//     get image() {
//         return this.#image;
//     }

//     // This is an abstract method that subclasses will implement
//     get author() {
//         throw new Error('getAuthor() method must be implemented');
//     }

//     // This is a polymorphic method that can take any subclass of Testimonial
//     get testimonialHTML() {
//         return `<div class="testimonial">
//             <img src="${this.image}" class="profile-testimonial" />
//             <p class="quote">"${this.quote}"</p>
//             <p class="author">- ${this.author}</p>
//         </div>`;
//     }
// }

// class AuthorTestimonial extends Testimonial {
//     #author = ""

//     constructor(author, quote, image) {
//         super(quote, image);
//         this.#author = author;
//     }

//     get author() {
//         return this.#author;
//     }
// }

// class CompanyTestimonial extends Testimonial {
//     #company = ""

//     constructor(company, quote, image) {
//         super(quote, image);
//         this.#company = company;
//     }

//     get author() {
//         return this.#company + " Company";
//     }
// }

// const testimonial1 = new AuthorTestimonial('Surya Elidanto', 'Mantap sekali jasanya!', 'https://images.unsplash.com/photo-1570295999919-56ceb5ecca61?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60');

// const testimonial2 = new AuthorTestimonial('Surya Elz', 'Keren lah pokoknya!', 'https://images.unsplash.com/photo-1568602471122-7832951cc4c5?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60');

// const testimonial3 = new CompanyTestimonial('ABC', 'Wuhuu keren lah!', 'https://images.unsplash.com/photo-1564564321837-a57b7070ac4f?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8OHx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60');

// let testimonialData = [testimonial1, testimonial2, testimonial3];
// let testimonialHTML = '';

// for (let i = 0; i < testimonialData.length; i++) {
//     testimonialHTML += testimonialData[i].testimonialHTML;
// }

// document.getElementById('testimonials').innerHTML = testimonialHTML;

// array of objects
const testimonialData = [
    {
        author: "Surya Elidanto",
        quote: "Keren banget jasanya!",
        image: "https://images.unsplash.com/photo-1570295999919-56ceb5ecca61?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
        rating: 5
    },
    {
        author: "Surya Elz",
        quote: "Keren lah pokoknya!",
        image: "https://images.unsplash.com/photo-1568602471122-7832951cc4c5?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
        rating: 4
    },
    {
        author: "Surya Gans",
        quote: "The best pelayanannya!",
        image: "https://images.unsplash.com/photo-1564564321837-a57b7070ac4f?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8OHx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
        rating: 4
    }
]

function allTestimonials() {
    let testimonialHTML = '';

    testimonialData.forEach(function (item) {
        testimonialHTML += `<div class="testimonial">
            <img src="${item.image}" class="profile-testimonial" />
            <p class="quote">"${item.quote}"</p>
            <p class="author">- ${item.author}</p>
            <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
        </div>`
    })

    document.getElementById('testimonials').innerHTML = testimonialHTML;
}

allTestimonials()


// filter testimonials
function filterTestimonials(rating) {
    let testimonialHTML = '';

    const testimonialFiltered = testimonialData.filter(function (item) {
        return item.rating === rating;
    })

    // console.log(testimonialFiltered);

    if (testimonialFiltered.length === 0) {
        testimonialHTML = `<h1> Data not found! </h1>`
    } else {
        testimonialFiltered.forEach(function (item) {
            testimonialHTML += `<div class="testimonial">
                <img src="${item.image}" class="profile-testimonial" />
                <p class="quote">"${item.quote}"</p>
                <p class="author">- ${item.author}</p>
                <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
            </div>`
        })
    }

    document.getElementById('testimonials').innerHTML = testimonialHTML;
}