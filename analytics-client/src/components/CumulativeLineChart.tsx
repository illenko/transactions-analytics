import {FC} from "react";
import {LineChart} from "@mui/x-charts";
import {Analytics} from "../App.types.ts";
import {Value} from "../constants/AnalyticsDatesEnums.ts";

interface ChartProps {
    analytics: Analytics | undefined;
    value: Value;
}

const CumulativeLineChart: FC<ChartProps> = ({analytics, value}) => {
    if (!analytics) return null;
    return (
        <LineChart
            width={1000}
            height={500}
            series={[
                {
                    data: analytics.groups.map(it => {
                        return value == Value.Amount ? Math.abs(it.amount) : it.count
                    }),
                }
            ]}
            xAxis={[{
                scaleType: 'point', data: analytics.groups.map(it => {
                    return it.name
                })
            }]}
        />
    );
};

export default CumulativeLineChart;