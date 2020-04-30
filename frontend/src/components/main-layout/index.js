import React from "react";
import PropTypes from 'prop-types'

import Header from '../header'

export default function MainLayout({ children }) {
    return (
        <>
            <Header></Header>
            <main className="container">{children}</main>
        </>
    )
}

MainLayout.propTypes = {
    children: PropTypes.node.isRequired
}