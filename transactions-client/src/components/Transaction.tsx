import {FC} from "react";
import {TransactionEntity} from "../App.types.ts";


const Transaction: FC<TransactionEntity> = ({id, datetime, amount, category, merchant}) => {
    return (
        <li>
            <div>
                Id: {id}
            </div>
            <div>
                Datetime: {datetime}
            </div>
            <div>
                Amount: {amount}
            </div>
            <div>
                Category: {category}
            </div>
            <div>
                Merchant: {merchant}
            </div>
            <hr/>
        </li>
    );
};

export default Transaction;