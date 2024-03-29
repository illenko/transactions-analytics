import {FC, useEffect, useState} from "react";
import {AppProps, StatisticsResult} from "../App.types.ts";
import axios from "axios";
import Container from "@mui/material/Container";
import Stack from "@mui/material/Stack";
import Button from "@mui/material/Button";
import StatisticCard from "../components/StatisticCard.tsx";

const Statistics: FC<AppProps> = () => {
    const [statistics, setStatistics] = useState<StatisticsResult>();

    const showStatisticsByCategory = async () => {
        try {
            const {data} = await axios.get('http://localhost:8080/statistics/category');
            setStatistics(data);
        } catch (error) {
            console.error(error);
        }
    };
    useEffect(() => {
        showStatisticsByCategory();
    }, []);
    

    const showStatisticsByMerchant = async () => {
        try {
            const {data} = await axios.get('http://localhost:8080/statistics/merchant');
            setStatistics(data);
        } catch (error) {
            console.log(error);
        }
    };

    return (
        <Container maxWidth="lg">
            <Stack alignItems="center" justifyContent="center" spacing={2} direction="row">
                <Button variant="outlined" onClick={showStatisticsByCategory}>By category</Button>
                <Button variant="outlined" onClick={showStatisticsByMerchant}>By merchant</Button>
            </Stack>
            <Stack justifyContent="center" paddingTop={2} spacing={2} direction="row">
                {(() => {
                    if (statistics) {
                        return <StatisticCard title={"Expenses"} statistics={statistics.expenses}
                                              by={"categories"}></StatisticCard>
                    }
                })()}
                {(() => {
                    if (statistics) {
                        return <StatisticCard title={"Income"} statistics={statistics.income}
                                              by={"categories"}></StatisticCard>
                    }
                })()}
            </Stack>
        </Container>
    );
};

export default Statistics;