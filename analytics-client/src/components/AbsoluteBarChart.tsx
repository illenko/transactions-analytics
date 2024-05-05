import {FC} from "react";
import {BarChart} from "@mui/x-charts";
import {Analytics} from "../App.types.ts";
import {Value} from "../constants/AnalyticsDatesEnums.ts";

interface ChartProps {
    analytics: Analytics | undefined;
    value: Value;
}

const AbsoluteBarChart: FC<ChartProps> = ({analytics, value}) => {
    if (!analytics) return null;
    return (
        <BarChart
            xAxis={[{
                scaleType: 'band', data: analytics.groups.map((it) => {
                    return it.name
                })
            }]}
            series={[{
                data: analytics.groups.map((it) => {
                    return value == Value.Amount ? Math.abs(it.amount) : it.count
                })
            }]}
            width={1000}
            height={500}
        />
    );
};

export default AbsoluteBarChart;