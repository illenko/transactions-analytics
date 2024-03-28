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

export interface StatisticsResult {
    income: Statistics;
    expenses: Statistics;
}

export interface Statistics {
    count: number;
    amount: number;
    groups: StatisticsGroup[];
}

export interface StatisticsGroup {
    name: string;
    count: number;
    amount: number;
}

export interface StatProps {
    title: string;
    statistics: Statistics;
    by: string;
}