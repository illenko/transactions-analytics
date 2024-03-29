import {FC, useEffect, useState} from "react";
import {MonthExpense, TransactionEntity} from "../App.types";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import PaymentIcon from "@mui/icons-material/Payment";
import {BarChart} from "@mui/x-charts";
import axios from "axios";

const chartSetting = {
    width: 350,
    height: 300
};
const valueFormatter = (value: number | null) => `${value} $`;

const TransactionCard: FC<TransactionEntity> = ({id, datetime, amount, category, merchant}) => {
    const [expenses, setExpenses] = useState<MonthExpense[]>([]);

    useEffect(() => {
        const getExpenses = async () => {
            try {
                const {data} = await axios.get(`http://localhost:8080/merchants/${merchant}/expenses`);
                setExpenses(data);
            } catch (error) {
                console.error(error);
            }
        };
        if (amount < 0) {
            getExpenses();
        }

    }, []);


    return (
        <Card sx={{minWidth: 350, minHeight: 300}}>
            <CardContent>
                <Typography sx={{fontSize: 14}} color="text.secondary" gutterBottom>
                    <PaymentIcon/>
                </Typography>
                <Typography sx={{fontSize: 14}} color="text.secondary" gutterBottom>
                    {id}
                </Typography>
                <Typography component="div">
                    <b>Datetime:</b> {datetime}
                </Typography>
                <Typography component="div">
                    <b>Amount:</b> {amount} $
                </Typography>
                <Typography component="div">
                    <b>Category:</b> {category}
                </Typography>
                <Typography component="div">
                    <b>Merchant</b> {merchant}
                </Typography>

                {(() => {
                    if (expenses.length) {
                        return (
                            <>
                                <Typography component="div">
                                    You spent <b>{expenses.map(it => {
                                    return Math.abs(it.amount)
                                }).reduce((partialSum, a) => partialSum + a, 0)}</b> $ for last 6 months
                                </Typography>
                                <BarChart
                                    dataset={expenses.map(it => {
                                        return {amount: Math.abs(it.amount), month: it.month}
                                    })}
                                    xAxis={[{scaleType: 'band', dataKey: 'month'}]}
                                    series={[
                                        {dataKey: 'amount', label: 'Amount', valueFormatter},
                                    ]}
                                    {...chartSetting}
                                />
                            </>)
                    }
                })()}

            </CardContent>
        </Card>
    );
};

export default TransactionCard;