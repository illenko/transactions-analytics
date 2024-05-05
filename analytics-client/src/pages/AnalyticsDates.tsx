import {FC, useEffect, useState, useCallback} from "react";
import {Analytics} from "../App.types.ts";
import Container from "@mui/material/Container";
import Stack from "@mui/material/Stack";
import {Alert, CircularProgress, MenuItem, SelectChangeEvent} from "@mui/material";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import axios from "axios";
import SelectFormControl from "../components/SelectFormControl";
import CumulativeLineChart from "../components/CumulativeLineChart";
import AbsoluteBarChart from "../components/AbsoluteBarChart";
import {Direction, Unit, Calculation, Value} from "../constants/AnalyticsDatesEnums.ts";
import Box from "@mui/material/Box";

const AnalyticsDates: FC = () => {
    const [direction, setDirection] = useState<Direction>(Direction.Expenses);
    const [unit, setUnit] = useState<Unit>(Unit.Month);
    const [calculation, setCalculation] = useState<Calculation>(Calculation.Cumulative);
    const [value, setValue] = useState<Value>(Value.Amount);
    const [analytics, setAnalytics] = useState<Analytics>()
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(false);

    const handleDirectionChange = (event: SelectChangeEvent) => {
        setDirection(event.target.value as Direction);
        loadAnalytics().catch(error => console.error(error));
    };

    const handleUnitChange = (event: SelectChangeEvent) => {
        setUnit(event.target.value as Unit);
        loadAnalytics().catch(error => console.error(error));
    };

    const handleCalculationChange = (event: SelectChangeEvent) => {
        setCalculation(event.target.value as Calculation);
        loadAnalytics().catch(error => console.error(error));
    };

    const handleValueChange = (event: SelectChangeEvent) => {
        setValue(event.target.value as Value);
        loadAnalytics().catch(error => console.error(error));
    };

    const loadAnalytics = useCallback(async () => {
        setLoading(true);
        setError(false);
        try {
            const {data} = await axios.get(`http://localhost:8080/analytics/${direction}/dates?unit=${unit}&calculation=${calculation}`);
            setAnalytics(data);
        } catch (error) {
            setError(true);
        } finally {
            setLoading(false);
        }
    }, [direction, unit, calculation]);

    useEffect(() => {
        loadAnalytics().catch(error => console.error(error));
    }, [loadAnalytics, value]);

    if (loading) {
        return (
            <Box
                display="flex"
                justifyContent="center"
                alignItems="center"
                height="100vh"
            >
                <CircularProgress/>
            </Box>
        );
    }

    if (error) {
        return <Alert severity="error">An error occurred...</Alert>;
    }

    return (
        <Container maxWidth="lg">
            <Stack alignItems="center" justifyContent="center" spacing={2} direction="row">
                <SelectFormControl labelId="direction-label" id="direction" value={direction}
                                   onChange={handleDirectionChange}>
                    <MenuItem value={Direction.Expenses}>Expenses</MenuItem>
                    <MenuItem value={Direction.Income}>Income</MenuItem>
                </SelectFormControl>
                <SelectFormControl labelId="unit-label" id="unit" value={unit} onChange={handleUnitChange}>
                    <MenuItem value={Unit.Month}>Month</MenuItem>
                    <MenuItem value={Unit.Day}>Day</MenuItem>
                </SelectFormControl>
                <SelectFormControl labelId="calculation-label" id="calculation" value={calculation}
                                   onChange={handleCalculationChange}>
                    <MenuItem value={Calculation.Cumulative}>Cumulative</MenuItem>
                    <MenuItem value={Calculation.Absolute}>Absolute</MenuItem>
                </SelectFormControl>
                <SelectFormControl labelId="value-label" id="value" value={value} onChange={handleValueChange}>
                    <MenuItem value={Value.Amount}>Amount</MenuItem>
                    <MenuItem value={Value.Count}>Count</MenuItem>
                </SelectFormControl>
            </Stack>

            <Stack justifyContent="center" paddingTop={2} spacing={2} direction="row">
                <Card sx={{minWidth: 500, minHeight: 300}}>
                    <CardContent>
                        {calculation === Calculation.Cumulative ? (
                            <CumulativeLineChart analytics={analytics} value={value}/>
                        ) : (
                            <AbsoluteBarChart analytics={analytics} value={value}/>
                        )}
                    </CardContent>
                </Card>
            </Stack>
        </Container>
    );
};

export default AnalyticsDates;