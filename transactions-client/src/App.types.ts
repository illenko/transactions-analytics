export interface AppProps {
    title: string;
}

export interface TransactionEntity {
    id: string,
    datetime: string,
    amount: string,
    category: string,
    merchant: string
}