import {Link} from "react-router-dom";

function Home(){
    return (
        <>
            <h1>H-vistars</h1>
            <Link to="/login">Login</Link>
            <br/>
            <Link to="/register">Register</Link>
        </>
    )
}

export default Home;