import {FC, useEffect, useState} from "react";
import {Analytic} from "../App.types.ts";
import Container from "@mui/material/Container";
import Stack from "@mui/material/Stack";
import {FormControl, InputLabel, MenuItem, Select, SelectChangeEvent} from "@mui/material";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import axios from "axios";
import {BarChart, LineChart} from "@mui/x-charts";

const AnalyticDates: FC = () => {
    const [direction, setDirection] = useState<string>('expenses');
    const [unit, setUnit] = useState<string>('month');
    const [calculation, setCalculation] = useState<string>('cumulative');
    const [value, setValue] = useState<string>('amount');
    const [analytic, setAnalytic] = useState<Analytic>()

    const handleDirectionChange = (event: SelectChangeEvent) => {
        setDirection(event.target.value);
        loadAnalytic()
    };

    const handleUnitChange = (event: SelectChangeEvent) => {
        setUnit(event.target.value);
        loadAnalytic()
    };

    const handleCalculationChange = (event: SelectChangeEvent) => {
        setCalculation(event.target.value);
        loadAnalytic()
    };

    const handleValueChange = (event: SelectChangeEvent) => {
        setValue(event.target.value);
        loadAnalytic()
    };

    const loadAnalytic = async () => {
        try {
            const {data} = await axios.get(`http://localhost:8080/analytic/${direction}/dates?unit=${unit}&calculation=${calculation}`);
            setAnalytic(data);
        } catch (error) {
            console.error(error);
        }
    };

    useEffect(() => {
        loadAnalytic();
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
                                if (analytic) {
                                    if (calculation == 'cumulative') {
                                        return <LineChart
                                            width={1000}
                                            height={500}
                                            series={[
                                                {
                                                    data: analytic.groups.map(it => {
                                                        return value == 'amount' ? Math.abs(it.amount) : it.count
                                                    }),
                                                }
                                            ]}
                                            xAxis={[{
                                                scaleType: 'point', data: analytic.groups.map(it => {
                                                    return it.name
                                                })
                                            }]}
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

export default AnalyticDates;