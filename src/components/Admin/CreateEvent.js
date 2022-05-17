// import React, { useState } from 'react';
import React from "react";
// import { useHistory } from 'react-router-dom';
// import axios from 'axios';
import clsx from "clsx";

import Header from '../Header/Header';

import { Button, Box, Card, Divider, Grid, Paper, Typography, TextField } from "@material-ui/core";

import { makeStyles } from '@material-ui/core/styles';

// import { ToastContainer, toast } from 'react-toastify';
// import 'react-toastify/dist/ReactToastify.css';

const useStyles = makeStyles((theme) => ({
    root: {
        backgroundColor: 'var(--lightgray-color)',
        position: 'absolute',
        display: 'flex',
        height: '100%',
        overflow: 'hidden',
        width: '100%',
    },
    wrapper: {
        display: 'flex',
        flex: '1 1 auto',
        overflow: 'hidden',
        paddingTop: 64,
        [theme.breakpoints.up('md')]: {
            paddingLeft: 256,
        },
        [theme.breakpoints.up('lg')]: {
            paddingLeft: 256,
        },
    },
    contentContainer: {
        display: 'flex',
        flex: '1 1 auto',
        overflow: 'hidden',
    },
    content: {
        flex: '1 1 auto',
        height: '100%',
        overflow: 'auto',
    },
    pageContainer: {
        minHeight: "100%",
        marginTop: "0px",
    },
    secondRoot: {
        margin: "20px 20px",
        marginBottom: "50px",
        padding: "20px 20px",
    },
    mainTable: {
        margin: "20px 0 20px",
    },
    title: {
        color: "var(--blue-color)",
        textAlign: "center",
        margin: "30px",
    },
    divider: {
        backgroundColor: "var(--blue-color)",
        height: "5px",
        margin: "10px 0",
    },
    editForm: {
        display: "grid",
        padding: "2% 5%",
    },
    inputField: {
        margin: "10px 0",
    },
    inputCont: {
        [theme.breakpoints.down("xs")]: {
            display: "block",
        },
    },
    btnContainer: {
        textAlign: "right",
        marginTop: '20px',
        [theme.breakpoints.down("xs")]: {
            textAlign: "center",
            display: "flex",
            justifyContent: "center",
        },
    },
    button: {
        background: "var(--blue-color)",
        color: "white",
        marginLeft: "30px",
        width: "100px",
        "&:hover": {
            backgroundColor: "#113e78ec",
        },
        [theme.breakpoints.down("xs")]: {
            margin: "15px",
        },
    },
}));

const CreateEvent = ({ className, staticContext, ...rest }) => {
    // const [customer, setCustomer] = useState({});
    // const history = useHistory();
    const classes = useStyles();

    // const submitCustomer = async (e) => {
    //     console.log(customer);
    //     e.preventDefault();
    //     const url = '';
    //     try {
    //         await axios.post(url, customer);
    //         toasterSucc();
    //     } catch (err) {
    //         err && toasterErr(err);
    //     }
    // };

    // const toasterSucc = () => {
    //     return (
    //         toast.success('Customer successfully created!', {
    //             position: "bottom-center",
    //             autoClose: 3000,
    //             onClose: () => history.push(`/admin/customers`),
    //             hideProgressBar: false,
    //             closeOnClick: true,
    //             pauseOnHover: true,
    //             draggable: true,
    //             progress: undefined
    //         })
    //     );
    // };

    // const toasterErr = (error) => {
    //     return (
    //         toast.error(`${error}`, {
    //             position: "bottom-center",
    //             autoClose: 3000,
    //             hideProgressBar: false,
    //             closeOnClick: true,
    //             pauseOnHover: true,
    //             draggable: true,
    //             progress: undefined
    //         })
    //     );
    // };

    return (
        <div className={classes.root}>
            <div className={classes.wrapper}>
                <div className={classes.contentContainer}>
                    <div className={classes.content}>

                        <Header />

                        <div className={clsx(classes.pageContainer, className)}>
                            <Card className={clsx(classes.secondRoot, className)} elevation={2} {...rest}>
                                <Typography className={classes.title} gutterBottom variant="h4" component="h1">
                                    Create a new event
                                </Typography>

                                <Divider className={classes.divider} />

                                <Grid>
                                    <Paper className={classes.mainTable}>
                                        <form
                                            // onSubmit={(e) => submitCustomer(e)} autoComplete="off" 
                                            className={classes.editForm}>
                                            <TextField
                                                className={classes.inputField}
                                                label="Title"
                                                type="text"
                                                name="title"
                                                placeholder="Enter a title"
                                                variant="outlined"
                                                fullWidth
                                                InputLabelProps={{ shrink: true }}
                                            // onChange={(e) => setCustomer({ ...customer, title: e.target.value })}
                                            />

                                            <TextField
                                                className={classes.inputField}
                                                label="Type"
                                                type="text"
                                                name="type"
                                                placeholder="Enter a type"
                                                variant="outlined"
                                                fullWidth
                                                InputLabelProps={{ shrink: true }}
                                            // onChange={(e) => setCustomer({ ...customer, type: e.target.value })}
                                            />

                                            <TextField
                                                className={classes.inputField}
                                                label="Location"
                                                type="text"
                                                name="location"
                                                placeholder="Enter a location"
                                                variant="outlined"
                                                fullWidth
                                                InputLabelProps={{ shrink: true }}
                                            // onChange={(e) => setCustomer({ ...customer, location: e.target.value })}
                                            />

                                            <Grid className={classes.inputCont} container justifyContent="space-between">
                                                <Box flex={1} mr={{ xs: 0, sm: "0.5em" }}>
                                                    <TextField
                                                        className={classes.inputField}
                                                        label="Start"
                                                        type="date"
                                                        name="start"
                                                        placeholder="Enter a start date"
                                                        variant="outlined"
                                                        fullWidth
                                                        InputLabelProps={{ shrink: true }}
                                                    // onChange={(e) => setCustomer({ ...customer, start: e.target.value })}
                                                    />
                                                </Box>

                                                <Box flex={1} ml={{ xs: 0, sm: "0.5em" }}>
                                                    <TextField
                                                        className={classes.inputField}
                                                        label="End"
                                                        type="date"
                                                        name="end"
                                                        placeholder="Enter a end date"
                                                        variant="outlined"
                                                        fullWidth
                                                        InputLabelProps={{ shrink: true }}
                                                    // onChange={(e) => setCustomer({ ...customer, end: e.target.value })}
                                                    />
                                                </Box>
                                            </Grid>
                                            <div className={classes.btnContainer}>
                                                <Button className={classes.button} href='/admin/events'>Back</Button>
                                                <Button className={classes.button} type="submit">Save</Button>
                                            </div>
                                        </form>
                                    </Paper>
                                </Grid>
                            </Card>
                            {/* <ToastContainer /> */}
                        </div>
                    </div >
                </div >
            </div >
        </div >
    );
}

export default CreateEvent;