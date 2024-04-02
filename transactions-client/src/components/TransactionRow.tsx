import {FC} from "react";
import {Transaction} from "../App.types.ts";
import TableCell from "@mui/material/TableCell";
import TableRow from "@mui/material/TableRow";


const TransactionRow: FC<Transaction> = ({id, datetime, amount, category, merchant}) => {
    return (
        <TableRow
            key={id}
            sx={{'&:last-child td, &:last-child th': {border: 0}}}>
            <TableCell component="th" scope="row">
                {id}
            </TableCell>
            <TableCell>{datetime}</TableCell>
            <TableCell>{amount}</TableCell>
            <TableCell>{category}</TableCell>
            <TableCell>{merchant}</TableCell>
        </TableRow>
    );
};

export default TransactionRow;