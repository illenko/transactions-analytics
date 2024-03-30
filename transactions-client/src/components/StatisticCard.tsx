import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import {FC} from "react";
import {StatProps} from "../App.types.ts";
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import PaymentIcon from '@mui/icons-material/Payment';
import {LineChart, PieChart} from "@mui/x-charts";

const StatisticCard: FC<StatProps> = ({title, statistics, by}) => {
    return (
        <Card sx={{minWidth: 275, minHeight: 300}}>
            <CardContent>
                <Typography sx={{fontSize: 14}} color="text.secondary" gutterBottom>
                    {title}
                </Typography>
                <Typography component="div">
                    Amount: <b>{statistics.amount}$</b>
                </Typography>
                <Typography component="div">
                    You have made <b>{statistics.count}</b> transactions
                </Typography>
                <Typography component="div">
                    in <b>{statistics.groups.length}</b> {by}
                </Typography>
                <PieChart
                    series={[
                        {
                            data: statistics.groups.map((it, index) => {
                                return {key: index, value: Math.abs(it.amount), label: it.name}
                            }),
                            innerRadius: 80,
                            outerRadius: 100,
                            paddingAngle: 1,
                            cornerRadius: 3,
                            cx: 150,
                            cy: 150,
                        },
                    ]}
                    width={400}
                    height={300}
                />
                <LineChart
                    width={500}
                    height={300}
                    series={[
                        {
                            data: statistics.dateAmounts.map(it => {
                                return Math.abs(it.amount)
                            })
                        },
                    ]}
                    xAxis={[{
                        scaleType: 'point', data: statistics.dateAmounts.map(it => {
                            return it.date + ''
                        })
                    }]}
                    sx={{
                        '.MuiLineElement-root, .MuiMarkElement-root': {
                            strokeWidth: 1,
                        },
                    }}
                />
                <Typography sx={{fontSize: 14}} component="div">
                    <List sx={{width: '100%', maxWidth: 360, bgcolor: 'background.paper'}}>
                        {statistics.groups.map(({name, amount, count}) => {
                            return <ListItem>
                                <ListItemAvatar>
                                    <PaymentIcon/>
                                </ListItemAvatar>
                                <ListItemText primary={name}
                                              secondary={count + " transaction" + (count == 1 ? "" : "s") + ", " + amount + "$"}/>
                            </ListItem>;
                        })}
                    </List>
                </Typography>
            </CardContent>
        </Card>
    );
};

export default StatisticCard;