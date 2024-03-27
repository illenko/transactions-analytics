import axios from 'axios';
import {FC, useState} from 'react';
import Transaction from "./components/Transaction.tsx";
import {AppProps, TransactionEntity} from "./App.types.ts";

const App: FC<AppProps> = ({title}) => {
    const [transactions, setTransactions] = useState<TransactionEntity[]>([]);

    const showTransactions = async () => {
        try {
            const {data} = await axios.get('http://localhost:8080/transactions');
            setTransactions(data);
        } catch (error) {
            console.log(error);
        }
    };
    
    return (
        <div>
            <h1>{title}</h1>
            <div className="centered">
                <button onClick={showTransactions}>Transactions</button>
            </div>

            <ul>
                {transactions.map(({id, datetime, amount, category, merchant}) => {
                    return <Transaction id={id} datetime={datetime} amount={amount} category={category}
                                        merchant={merchant}/>;
                })}
            </ul>
        </div>
    );
};

export default App;