import Navbar from "../components/navbar";
import './home.css';
import Search from "../components/search";

export default function Home() {
    return (
        <>
        <Navbar />
        <div>
            <h1>op.gg</h1>
            <Search />
        </div>
        </>
    )
}