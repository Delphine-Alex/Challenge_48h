import React, { Fragment, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import PropTypes from 'prop-types';

import { Box, Drawer, Hidden, List, Typography, makeStyles } from '@material-ui/core';

import CalendarMonthIcon from '@mui/icons-material/CalendarMonth';
import GroupIcon from '@mui/icons-material/Group';
import PersonIcon from '@mui/icons-material/Person';
import LogoutIcon from '@mui/icons-material/Logout';

import NavItem from './NavItem';

const items = [
    {
        href: '/admin/customers',
        icon: PersonIcon,
        title: 'User list'
    },
    {

        href: '/admin/organizers',
        icon: GroupIcon,
        title: 'Organizer list',
    },
    {
        href: '/admin/events',
        icon: CalendarMonthIcon,
        title: 'Event list',
    }
];

const useStyles = makeStyles(() => ({
    mobileDrawer: {
        backgroundColor: 'var(--blue-color)',
        width: 256
    },
    desktopDrawer: {
        backgroundColor: 'var(--blue-color)',
        width: 256,
        top: 64,
        height: 'calc(100% - 64px)'
    },
    name: {
        marginTop: '15px',
        color: 'white',
        textAlign: 'center'
    },
}));

const Navbar = ({ onMobileClose, openMobile }) => {
    const classes = useStyles();
    const location = useLocation();

    useEffect(() => {
        if (openMobile && onMobileClose === false) {
            onMobileClose();
        }
    }, [location.pathname, onMobileClose, openMobile]);

    const content = (
        <Box height="100%" display="flex" flexDirection="column" >
            <Box alignItems="center" display="flex" flexDirection="column" p={2} >
                <Typography className={classes.name} variant="h6">
                    Menu
                </Typography>
            </Box>
            <Box p={2}>
                <List>
                    {items.map((item) => (
                        <NavItem
                            href={item.href}
                            key={item.title}
                            title={item.title}
                            icon={item.icon}
                        />
                    ))}
                </List>
            </Box>
            <Box flexGrow={1} />
            <Box p={2} m={2} >
                <List>
                    <NavItem
                        href='/admin'
                        key='Logout'
                        title='Logout'
                        icon={LogoutIcon}
                    />
                </List>
            </Box>
        </Box>
    );

    return (
        <Fragment>
            <Hidden mdUp>
                <Drawer
                    anchor="left"
                    classes={{ paper: classes.mobileDrawer }}
                    onClose={onMobileClose}
                    open={openMobile}
                    variant="temporary"
                >
                    {content}
                </Drawer>
            </Hidden>
            <Hidden smDown>
                <Drawer
                    anchor="left"
                    classes={{ paper: classes.desktopDrawer }}
                    open
                    variant="persistent"
                >
                    {content}
                </Drawer>
            </Hidden>
        </Fragment>
    );
};

Navbar.propTypes = {
    onMobileClose: PropTypes.func,
    openMobile: PropTypes.bool
};

Navbar.defaultProps = {
    onMobileClose: () => { },
    openMobile: false
};

export default Navbar;