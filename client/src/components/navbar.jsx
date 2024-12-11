import React from 'react'
import { Link } from 'react-router-dom'
import githubLogo from '../icons/github-mark.svg'
import './navbar.css'

const Navbar = () => {
    return (
        <div className='nav-container'>
            <nav className='navbar'>
                <div className='main-content'>
                    <ul className='route-list'>
                        <li className='route-item'><Link to="/">Home</Link></li>
                        <li className='route-item'><Link to="/">Champions</Link></li>
                        <li className='route-item'><Link to="/">Game modes</Link></li>
                        <li className='route-item'><Link to="/">Leaderboards</Link></li>
                        <li className='route-item'><Link to="/">Pro spectate</Link></li>
                        <li className='route-item'><Link to="/">Stats</Link></li>
                        <li className='route-item'><Link to="/">Multi-search</Link></li>
                    </ul>
                </div>
                <div className='right-content'>
                    <ul className='route-list'>
                        <li className='route-item'>
                            <Link to="https://github.com/jonathanjchung">
                                <img src={githubLogo} alt="Github" width="25" height="25" />
                            </Link>
                        </li>
                    </ul>
                </div>
            </nav>
        </div>
    );
};

export default Navbar;