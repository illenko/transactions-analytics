import {FC, useEffect, useState} from "react";
import {Analytic} from "../App.types.ts";
import Container from "@mui/material/Container";
import Stack from "@mui/material/Stack";
import {FormControl, InputLabel, MenuItem, Select, SelectChangeEvent} from "@mui/material";
import Typography from "@mui/material/Typography";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import axios from "axios";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemAvatar from "@mui/material/ListItemAvatar";
import PaymentIcon from "@mui/icons-material/Payment";
import ListItemText from "@mui/material/ListItemText";
import {BarChart, PieChart} from "@mui/x-charts";

const AnalyticGroups: FC = () => {
    const [direction, setDirection] = useState<string>('expenses');
    const [group, setGroup] = useState<string>('category');
    const [chartType, setChartType] = useState<string>('pie');
    const [analytic, setAnalytic] = useState<Analytic>()

    const handleDirectionChange = (event: SelectChangeEvent) => {
        setDirection(event.target.value);
        loadAnalytic()
    };

    const handleGroupChange = (event: SelectChangeEvent) => {
        setGroup(event.target.value);
        loadAnalytic()
    };

    const handleChartTypeChange = (event: SelectChangeEvent) => {
        setChartType(event.target.value);
        loadAnalytic()
    };

    const loadAnalytic = async () => {
        try {
            const {data} = await axios.get(`http://localhost:8080/analytic/${direction}/groups?group=${group}`);
            setAnalytic(data);
        } catch (error) {
            console.error(error);
        }
    };

    useEffect(() => {
        loadAnalytic();
    }, [direction, group]);


    return (
        <>
            <Container maxWidth="lg">
                <Stack alignItems="center" justifyContent="center" spacing={2} direction="row">
                    <FormControl variant="filled" sx={{m: 1, minWidth: 300}}>
                        <InputLabel id="direction-label">Direction</InputLabel>
                        <Select labelId="direction-label" id="type" value={direction} onChange={handleDirectionChange}>
                            <MenuItem value="expenses">Expenses</MenuItem>
                            <MenuItem value="income">Income</MenuItem>
                        </Select>
                    </FormControl>
                    <FormControl variant="filled" sx={{m: 1, minWidth: 300}}>
                        <InputLabel id="group-label">Group by</InputLabel>
                        <Select labelId="group-label" id="group" value={group} onChange={handleGroupChange}>
                            <MenuItem value="category">Category</MenuItem>
                            <MenuItem value="merchant">Merchant</MenuItem>
                        </Select>
                    </FormControl>
                    <FormControl variant="filled" sx={{m: 1, minWidth: 300}}>
                        <InputLabel id="chart-type-label">Chart type</InputLabel>
                        <Select labelId="chart-type-label" id="chart-type" value={chartType}
                                onChange={handleChartTypeChange}>
                            <MenuItem value="pie">Pie</MenuItem>
                            <MenuItem value="bar">Bar</MenuItem>
                        </Select>
                    </FormControl>
                </Stack>

                <Stack justifyContent="center" paddingTop={2} spacing={2} direction="row">
                    <Card sx={{minWidth: 400, minHeight: 300}}>
                        <CardContent>
                            <Typography component="div">
                                Amount: <b>{analytic?.amount}</b> $
                            </Typography>
                            <Typography component="div">
                                You have made <b>{analytic?.count}</b> transactions
                            </Typography>
                            <Typography component="div">
                                in <b>{analytic?.groups?.length}</b> {group}
                            </Typography>
                            <Typography sx={{fontSize: 14}} component="div">
                                <List sx={{width: '100%', maxWidth: 360, bgcolor: 'background.paper'}}>
                                    {analytic?.groups.map(({name, amount, count}) => {
                                        return <ListItem>
                                            <ListItemAvatar>
                                                <PaymentIcon/>
                                            </ListItemAvatar>
                                            <ListItemText primary={name}
                                                          secondary={count + " transaction" + (count == 1 ? "" : "s") + ", " + amount + "$"}/>
                                        </ListItem>;
                                    })}
                                </List>
                            </Typography>
                        </CardContent>
                    </Card>
                    <Card sx={{minWidth: 500, minHeight: 300}}>
                        <CardContent>
                            {(() => {
                                if (analytic) {
                                    if (chartType == 'pie') {
                                        return <PieChart
                                            series={[
                                                {
                                                    data: analytic.groups.map((it, index) => {
                                                        return {key: index, value: Math.abs(it.amount), label: it.name}
                                                    }),
                                                    innerRadius: 80,
                                                    outerRadius: 100,
                                                    paddingAngle: 1,
                                                    cornerRadius: 3,
                                                    cx: 150,
                                                    cy: 150,
                                                },
                                            ]}
                                            width={700}
                                            height={300}
                                        />
                                    } else {
                                        return <BarChart
                                            xAxis={[{
                                                scaleType: 'band', data: analytic.groups.map((it) => {
                                                    return it.name
                                                })
                                            }]}
                                            series={[{
                                                data: analytic.groups.map((it) => {
                                                    return Math.abs(it.amount)
                                                })
                                            }]}
                                            width={700}
                                            height={300}
                                        />
                                    }
                                }
                            })()}
                        </CardContent>
                    </Card>
                </Stack>
            </Container>
        </>
    );
};

export default AnalyticGroups;