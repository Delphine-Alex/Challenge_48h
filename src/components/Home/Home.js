import React from 'react';

import { Button, Typography } from '@material-ui/core';

import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles((theme) => ({
    container: {
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        padding: '0px 50px 50px 50px',
        textAlign: 'center',
        [theme.breakpoints.down('xs')]: {
            flexDirection: 'column'
        },
    },
    subcontainer: {
        padding: 40
    },
    titlePage: {
        color: 'var(--blue-color)',
        fontWeight: 'bold',
        textAlign: 'center',
    },
    subtitlePage: {
        textAlign: 'center',
        fontSize: '1em',
    },
    title: {
        color: 'var(--blue-color)',
        fontWeight: 'bold',
        fontSize: '1.25em',
        [theme.breakpoints.down('md')]: {
            fontSize: '1em',
        },
    },
    subtitle: {
        color: 'var(--black-color)',
        fontSize: '1em',
        [theme.breakpoints.down('md')]: {
            fontSize: '0.75em',
        },
    },
    btnContainer: {
        display: 'flex',
        justifyContent: 'center',
    },
    joinBtn: {
        backgroundColor: 'var(--blue-color)',
        color: 'var(--white-color)',
    }
}));

const Home = (className, staticContext) => {
    const classes = useStyles();

    return (
        <div>
            <Typography className={classes.titlePage} gutterBottom variant="h4" component="h1">How Meetup works</Typography>
            <p className={classes.subtitlePage}>Meet new people who share your interests through online and in-person events. It’s free to create an account.</p>

            <div className={classes.container}>
                <div className={classes.subcontainer}>
                    <img src='../images/handsUp.png' />
                    <Typography className={classes.title} gutterBottom variant="h4" component="h2">Join a group</Typography>
                    <Typography className={classes.subtitle} variant="subtitle1" gutterBottom component="p">Do what you love, meet others who love it, find your community. The rest is history!</Typography>
                </div>
                <div className={classes.subcontainer}>
                    <img src='../images/joinGroup.png' />
                    <Typography className={classes.title} gutterBottom variant="h4" component="h2">Find a event</Typography>
                    <Typography className={classes.subtitle} variant="subtitle1" gutterBottom component="p">Events are happening on just about any topic you can think of, from online gaming and photography to yoga and hiking.</Typography>
                </div>
                <div className={classes.subcontainer}>
                    <img src='../images/ticket.png' />
                    <Typography className={classes.title} gutterBottom variant="h4" component="h2">Start a group</Typography>
                    <Typography className={classes.subtitle} variant="subtitle1" gutterBottom component="p">You don’t have to be an expert to gather people together and explore shared interests.</Typography>
                </div>
            </div>

            <div className={classes.btnContainer}>
                <Button className={classes.joinBtn}>Join Meetup Here</Button>
            </div>
        </div>
    );
}

export default Home;
