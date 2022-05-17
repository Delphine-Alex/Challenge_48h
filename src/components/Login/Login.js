import React, { useState } from 'react';
import clsx from 'clsx';

import { Button, Card, Container, makeStyles } from '@material-ui/core';

import InputLabel from '@material-ui/core/InputLabel';
import OutlinedInput from '@material-ui/core/OutlinedInput';
import InputAdornment from '@material-ui/core/InputAdornment';
import IconButton from '@material-ui/core/IconButton';
import Visibility from '@material-ui/icons/Visibility';
import VisibilityOff from '@material-ui/icons/VisibilityOff';
import FormControl from '@material-ui/core/FormControl';

const useStyles = makeStyles((theme) => ({
    root: {
        backgroundColor: 'var(--blue-color)',
        display: "flex",
        justifyContent: "center",
        width: '100%',
        height: '100%',
        position: 'absolute',
        top: '0',
        left: '0'
    },
    card: {
        textAlign: 'center',
        alignSelf: 'center',
        padding: '60px',
        [theme.breakpoints.down('xs')]: {
            padding: '15px',
            margin: '10px'
        }
    },
    title: {
        color: 'var(--blue-color)',
        fontSize: '1.5em',
    },
    cardContent: {
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        marginTop: 20,
        height: "100%"
    },
    formEl: {
        marginTop: 20,
    },
    text: {
        fontSize: '0.7em',
        color: 'var(--blue-color)',
        marginTop: '15px',
    },
    loginBtn: {
        backgroundColor: 'var(--blue-color)',
        "&:hover": {
            backgroundColor: 'var(--hover-btn)'
        },
        height: '50px',
        color: 'white',
        marginTop: 30,
        marginBottom: 20
    }
}));

const Login = ({ className, staticContext, ...rest }) => {
    const classes = useStyles();
    const [login, setLogin] = useState('');
    const [password, setPassword] = useState('');
    const [showPassword, setShowPassword] = useState(false);

    const handleClickShowPassword = () => {
        setShowPassword(!showPassword);
    };

    const handleMouseDownPassword = (event) => {
        event.preventDefault();
    };

    return (
        <div className={classes.root}>
            <Card className={clsx(classes.card, className)} elevation={2} {...rest} >
                <h1 className={classes.title}>Paris Ynov Campus</h1>
                <Container maxWidth="sm" style={{ padding: 0 }}>

                    <form className={classes.cardContent} autoComplete="off">
                        <FormControl className={classes.formEl} variant="outlined">
                            <InputLabel htmlFor="email">Username</InputLabel>
                            <OutlinedInput
                                id="email"
                                label="Username"
                                variant="outlined"
                                onChange={(e) => setLogin(e.target.value)}
                                value={login}
                            />
                        </FormControl>

                        <FormControl className={classes.formEl} variant="outlined">
                            <InputLabel htmlFor="password">Password</InputLabel>
                            <OutlinedInput
                                id="password"
                                label="Password"
                                variant="outlined"
                                type={showPassword ? 'text' : 'password'}
                                onChange={(e) => setPassword(e.target.value)}
                                value={password}
                                endAdornment={
                                    <InputAdornment position="end">
                                        <IconButton
                                            aria-label="toggle password visibility"
                                            onClick={handleClickShowPassword}
                                            onMouseDown={handleMouseDownPassword}
                                            edge="end"
                                        >
                                            {showPassword ? <Visibility /> : <VisibilityOff />}
                                        </IconButton>
                                    </InputAdornment>
                                }
                            />
                        </FormControl>

                        <Button className={classes.loginBtn} href='/'>
                            Log In
                        </Button>
                    </form>
                </Container>
            </Card>
        </div>
    )
}

export default Login;