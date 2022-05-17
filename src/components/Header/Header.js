import React, { useState } from 'react';

import PropTypes from 'prop-types';
import clsx from 'clsx';

import NavBar from './NavBar';

import { AppBar, Box, Hidden, IconButton, Toolbar, Typography, makeStyles } from '@material-ui/core';

import MenuIcon from '@mui/icons-material/Menu';

const useStyles = makeStyles((theme) => ({
    root: {
        backgroundColor: 'var(--white-color)',
    },
    toolbar: {
        height: 64,
        backgroundColor: 'var(--white-color)',
    },
    title: {
        color: 'var(--black-color)',
        [theme.breakpoints.down('xl')]: {
            marginLeft: -70,
        },
        [theme.breakpoints.down('lg')]: {
            marginLeft: 20,
        },
        [theme.breakpoints.down('md')]: {
            marginLeft: 70,
        },
        [theme.breakpoints.down('sm')]: {
            fontSize: 15,
        },
        [theme.breakpoints.down('xs')]: {
            fontSize: 12,
        },
    }
}));

const Headers = ({ className, ...rest }) => {
    const classes = useStyles();
    const [isMobileNavOpen, setMobileNavOpen] = useState(false);

    return (
        <AppBar className={clsx(classes.root, className)} elevation={0} {...rest}>
            <NavBar
                onMobileClose={() => setMobileNavOpen(false)}
                openMobile={isMobileNavOpen}
            />
            <Toolbar>
                <Typography className={classes.title} variant="h5" component="h1">
                    Administrator Dashboard
                </Typography>

                <Box flexGrow={8} />
                <Hidden mdUp>
                    <IconButton color="inherit" onClick={() => setMobileNavOpen(true)}>
                        <MenuIcon style={{ color: '#113D78' }} />
                    </IconButton>
                </Hidden>
            </Toolbar>
        </AppBar >
    );
};

Headers.propTypes = {
    className: PropTypes.string,
};
export default Headers;