import React from 'react'
import { Link } from 'react-router-dom'
import './navbar.css'

const Navbar = () => {
    return (
        <nav className='navbar'>
            <div className='main-content'>
                <ul>
                    <li><Link to="/">Home</Link></li>
                    <li><Link to="/">Champions</Link></li>
                    <li><Link to="/">Game modes</Link></li>
                    <li><Link to="/">Leaderboards</Link></li>
                    <li><Link to="/">Pro spectate</Link></li>
                    <li><Link to="/">Stats</Link></li>
                    <li><Link to="/">Multi-search</Link></li>
                </ul>
            </div>
            <div className='right-content'>
                <ul>
                    <li><Link to="https://github.com/jonathanjchung">Github</Link></li>
                </ul>
            </div>
        </nav>
    );
};

export default Navbar;