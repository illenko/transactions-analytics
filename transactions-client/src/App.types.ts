export interface Transaction {
    id: string,
    datetime: string,
    amount: number,
    category: string,
    merchant: string
}

export interface Analytic {
    count: number
    amount: number
    groups: AnalyticGroup[]
}

export interface AnalyticGroup {
    name: string
    count: number
    amount: number
}
