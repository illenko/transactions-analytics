import {FC, useEffect, useState} from "react";
import axios from "axios";
import Container from "@mui/material/Container";
import TransactionCard from "../components/TransactionCard.tsx";
import {AppProps, TransactionEntity} from "../App.types.ts";
import {useParams} from "react-router-dom";
import Stack from "@mui/material/Stack";

const Transaction: FC<AppProps> = () => {
    const [transaction, setTransaction] = useState<TransactionEntity>();
    const { id } = useParams();

    const loadTransaction = async () => {
        try {
            const {data} = await axios.get(`http://localhost:8080/transactions/${id}`);
            setTransaction(data);
        } catch (error) {
            console.error(error);
        }
    };
    useEffect(() => {
        loadTransaction();
    }, []);

    return (
        <Container maxWidth="lg">
            <Stack justifyContent="center" paddingTop={2} spacing={2} direction="row">
            {(() => {
                if (transaction) {
                    return <TransactionCard id={transaction.id} datetime={transaction.datetime}
                                            amount={transaction.amount} category={transaction.category}
                                            merchant={transaction.merchant}></TransactionCard>
                }
            })()}
            </Stack>
            
        </Container>
    );
};

export default Transaction;