import React, { useState } from "react";
import LoginForm from "../common/forms/LoginForm.jsx";
import RegisterForm from "../common/forms/RegisterForm.jsx";

function Authentication(props) {
    const [isLogin, setIsLogin] = useState(props.option === "login");

    function handleToggle() {
        setIsLogin((prev) => !prev);
    }

    return (
        <div>
            <h1>{isLogin ? "Login" : "Register"}</h1>
            {isLogin ? (
                <LoginForm onToggle={handleToggle} />
            ) : (
                <RegisterForm onToggle={handleToggle} />
            )}
        </div>
    );
}

export default Authentication;

