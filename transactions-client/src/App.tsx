import axios from 'axios';
import {FC, useState} from 'react';
import Transaction from "./components/Transaction.tsx";
import {AppProps, StatisticsResult, TransactionEntity} from "./App.types.ts";
import StatisticCard from "./components/StatisticCard.tsx";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';
import Paper from '@mui/material/Paper';
import Button from '@mui/material/Button';
import Stack from '@mui/material/Stack';
import Typography from '@mui/material/Typography';
import Container from '@mui/material/Container';

const App: FC<AppProps> = ({title}) => {
    const [transactions, setTransactions] = useState<TransactionEntity[]>([]);
    const [statistics, setStatistics] = useState<StatisticsResult>();

    const showTransactions = async () => {
        try {
            const {data} = await axios.get('http://localhost:8080/transactions');
            setTransactions(data);
            setStatistics(undefined)
        } catch (error) {
            console.log(error);
        }
    };

    const showStatisticsByCategory = async () => {
        try {
            const {data} = await axios.get('http://localhost:8080/statistics/category');
            setStatistics(data);
            setTransactions([])
        } catch (error) {
            console.log(error);
        }
    };

    const showStatisticsByMerchant = async () => {
        try {
            const {data} = await axios.get('http://localhost:8080/statistics/merchant');
            setStatistics(data);
            setTransactions([])
        } catch (error) {
            console.log(error);
        }
    };

    return (
        <Container maxWidth="md">
            <Stack alignItems="center" justifyContent="center" paddingBottom={1} direction="row">
                <Typography alignItems="center" justifyContent="center" variant="h3" component="h3">{title}</Typography>
            </Stack>
            <Stack alignItems="center" justifyContent="center" spacing={2} direction="row">
                <Button variant="outlined" onClick={showTransactions}>Transactions</Button>
                <Button variant="outlined" onClick={showStatisticsByCategory}>Statistics by category</Button>
                <Button variant="outlined" onClick={showStatisticsByMerchant}>Statistics by merchant</Button>
            </Stack>
            {(() => {
                if (transactions.length) {
                    return <TableContainer className={"centered"} component={Paper}>
                        <Table>
                            <TableHead>
                                <TableRow>
                                    <TableCell>Id</TableCell>
                                    <TableCell>Datetime</TableCell>
                                    <TableCell>Amount</TableCell>
                                    <TableCell>Category</TableCell>
                                    <TableCell>Merchant</TableCell>
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {transactions.map(({id, datetime, amount, category, merchant}) => {
                                    return <Transaction id={id} datetime={datetime} amount={amount} category={category}
                                                        merchant={merchant}/>;
                                })}
                            </TableBody>
                        </Table>
                    </TableContainer>
                }
            })()}
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

export default App;