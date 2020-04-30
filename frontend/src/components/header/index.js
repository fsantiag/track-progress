import React from 'react'

import { Link, useLocation } from 'react-router-dom'

import './style.scss'

export default function Header() {
    const pathname = useLocation()

    const navLinks = () => (
        <>
            <Link to="/home" className={`nav-item nav-link ${pathname === '/home' ? 'active' : ''}`}> Home </Link>
            <Link to="/tasks" className={`nav-item nav-link ${pathname === '/tasks' ? 'active' : ''}`}> Tasks </Link>
        </>
    )

    return (
        <nav className="navbar navbar-expand-lg navbar-dark bg-dark">
            <span className="logo-name">Track Progress</span>
            <svg className="navbar-brand bi bi-check-all" width="2em" height="2em" viewBox="0 0 16 16" fill="currentColor" xmlns="http://www.w3.org/2000/svg">
                <path fillRule="evenodd" d="M12.354 3.646a.5.5 0 010 .708l-7 7a.5.5 0 01-.708 0l-3.5-3.5a.5.5 0 11.708-.708L5 10.293l6.646-6.647a.5.5 0 01.708 0z" clipRule="evenodd" />
                <path d="M6.25 8.043l-.896-.897a.5.5 0 10-.708.708l.897.896.707-.707zm1 2.414l.896.897a.5.5 0 00.708 0l7-7a.5.5 0 00-.708-.708L8.5 10.293l-.543-.543-.707.707z" />
            </svg>
            <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNavAltMarkup" aria-controls="navbarNavAltMarkup" aria-expanded="false" aria-label="Toggle navigation">
                <span className="navbar-toggler-icon"></span>
            </button>
            <div className="collapse navbar-collapse" id="navbarNavAltMarkup">
                <div className="navbar-nav">
                    {navLinks()}
                </div>
            </div>
        </nav>
    )
}