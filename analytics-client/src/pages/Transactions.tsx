import {FC, useEffect, useState} from "react";
import {Transaction} from "../App.types.ts";
import axios from "axios";
import Container from "@mui/material/Container";
import TableContainer from "@mui/material/TableContainer";
import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import TableCell from "@mui/material/TableCell";
import TableBody from "@mui/material/TableBody";
import TransactionRow from "../components/TransactionRow.tsx";

const Transactions: FC = () => {
    const [transactions, setTransactions] = useState<Transaction[]>([]);

    useEffect(() => {
        const getTransactions = async () => {
            try {
                const {data} = await axios.get('http://localhost:8080/transactions');
                setTransactions(data);
            } catch (error) {
                console.error(error);
            }
        };
        getTransactions();
    }, []);

    return (
        <Container maxWidth="lg">
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
                                    return <TransactionRow id={id} datetime={datetime} amount={amount}
                                                           category={category}
                                                           merchant={merchant}/>;
                                })}
                            </TableBody>
                        </Table>
                    </TableContainer>
                }
            })()}
        </Container>
    );
};

export default Transactions;