import {useState} from "react";

function RegisterForm() {
    const [formData, setFormData] = useState({
        username: '',
        email: '',
        password: '',
        country: '',
    });

    function handleInputChange(event) {
        const {name, value} = event.target;
        setFormData((prevFormData) => ({
            ...prevFormData,
            [name]: value
        }));
    }

    function handleSubmit(event) {
        event.preventDefault();
        fetch("localhost:8080/api/login", {
            method: "POST",
            body: JSON.stringify(formData),
            headers: {
                "Content-Type": "application/json"
            }
        })
            .then((response) => response.json())
            .then((data) => {
                console.log(data);
            })
            .catch((error) => {
                console.log(error);
            });
    }

    return (
        <div>
            <form onSubmit={handleSubmit}>
                <label htmlFor="username">Username</label>
                <input type="text" name="username" id="username" value={formData.username}
                       onChange={handleInputChange}/>
                <label htmlFor="email">Email</label>
                <input type="email" name="email" id="email" value={formData.email} onChange={handleInputChange}/>
                <label htmlFor="password">Password</label>
                <input type="password" name="password" id="password" value={formData.password}
                       onChange={handleInputChange}/>
                <label htmlFor="country">Country</label>
                <input type="text" name="country" id="country" value={formData.country}
                       onChange={handleInputChange}/>
                <button type="submit">Register</button>
            </form>
        </div>
    );
}

export default RegisterForm;