export interface Transaction {
    id: string,
    datetime: string,
    amount: number,
    category: string,
    merchant: string
}

export interface Analytics {
    count: number
    amount: number
    groups: AnalyticsGroup[]
}

export interface AnalyticsGroup {
    name: string
    count: number
    amount: number
}
