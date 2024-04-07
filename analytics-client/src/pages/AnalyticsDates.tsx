import {FC, useEffect, useState} from "react";
import {Analytics} from "../App.types.ts";
import Container from "@mui/material/Container";
import Stack from "@mui/material/Stack";
import {FormControl, InputLabel, MenuItem, Select, SelectChangeEvent} from "@mui/material";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import axios from "axios";
import {BarChart, LineChart} from "@mui/x-charts";

const AnalyticsDates: FC = () => {
    const [direction, setDirection] = useState<string>('expenses');
    const [unit, setUnit] = useState<string>('month');
    const [calculation, setCalculation] = useState<string>('cumulative');
    const [value, setValue] = useState<string>('amount');
    const [analytics, setAnalytics] = useState<Analytics>()

    const handleDirectionChange = (event: SelectChangeEvent) => {
        setDirection(event.target.value);
        loadAnalytics()
    };

    const handleUnitChange = (event: SelectChangeEvent) => {
        setUnit(event.target.value);
        loadAnalytics()
    };

    const handleCalculationChange = (event: SelectChangeEvent) => {
        setCalculation(event.target.value);
        loadAnalytics()
    };

    const handleValueChange = (event: SelectChangeEvent) => {
        setValue(event.target.value);
        loadAnalytics()
    };

    const loadAnalytics = async () => {
        try {
            const {data} = await axios.get(`http://localhost:8080/analytics/${direction}/dates?unit=${unit}&calculation=${calculation}`);
            setAnalytics(data);
        } catch (error) {
            console.error(error);
        }
    };

    useEffect(() => {
        loadAnalytics();
    }, [direction, unit, calculation, value]);


    return (
        <>
            <Container maxWidth="lg">
                <Stack alignItems="center" justifyContent="center" spacing={2} direction="row">
                    <FormControl variant="filled" sx={{m: 1, minWidth: 300}}>
                        <InputLabel id="direction-label">Type</InputLabel>
                        <Select labelId="direction-label" id="direction" value={direction}
                                onChange={handleDirectionChange}>
                            <MenuItem value="expenses">Expenses</MenuItem>
                            <MenuItem value="income">Income</MenuItem>
                        </Select>
                    </FormControl>
                    <FormControl variant="filled" sx={{m: 1, minWidth: 300}}>
                        <InputLabel id="unit-label">Unit</InputLabel>
                        <Select labelId="unit-label" id="unit" value={unit} onChange={handleUnitChange}>
                            <MenuItem value="month">Month</MenuItem>
                            <MenuItem value="day">Day</MenuItem>
                        </Select>
                    </FormControl>
                    <FormControl variant="filled" sx={{m: 1, minWidth: 300}}>
                        <InputLabel id="calculation-label">Calculation type</InputLabel>
                        <Select labelId="calculation-label" id="calculation" value={calculation}
                                onChange={handleCalculationChange}>
                            <MenuItem value="cumulative">Cumulative</MenuItem>
                            <MenuItem value="absolute">Absolute</MenuItem>
                        </Select>
                    </FormControl>
                    <FormControl variant="filled" sx={{m: 1, minWidth: 300}}>
                        <InputLabel id="value-label">Value</InputLabel>
                        <Select labelId="value-label" id="value" value={value}
                                onChange={handleValueChange}>
                            <MenuItem value="amount">Amount</MenuItem>
                            <MenuItem value="count">Count</MenuItem>
                        </Select>
                    </FormControl>
                </Stack>

                <Stack justifyContent="center" paddingTop={2} spacing={2} direction="row">

                    <Card sx={{minWidth: 500, minHeight: 300}}>
                        <CardContent>
                            {(() => {
                                if (analytics) {
                                    if (calculation == 'cumulative') {
                                        return <LineChart
                                            width={1000}
                                            height={500}
                                            series={[
                                                {
                                                    data: analytics.groups.map(it => {
                                                        return value == 'amount' ? Math.abs(it.amount) : it.count
                                                    }),
                                                }
                                            ]}
                                            xAxis={[{
                                                scaleType: 'point', data: analytics.groups.map(it => {
                                                    return it.name
                                                })
                                            }]}
                                        />
                                    } else {
                                        return <BarChart
                                            xAxis={[{
                                                scaleType: 'band', data: analytics.groups.map((it) => {
                                                    return it.name
                                                })
                                            }]}
                                            series={[{
                                                data: analytics.groups.map((it) => {
                                                    return value == 'amount' ? Math.abs(it.amount) : it.count
                                                })
                                            }]}
                                            width={1000}
                                            height={500}
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

export default AnalyticsDates;