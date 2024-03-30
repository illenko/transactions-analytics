export interface AppProps {
    title: string;
}

export interface TransactionEntity {
    id: string,
    datetime: string,
    amount: number,
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
    dateAmounts: DateAmount[];
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

export interface MonthAmount {
    month: string;
    amount: number;
}

export interface DateAmount {
    date: string;
    amount: number;
}