import {FC, ReactNode} from "react";
import {FormControl, InputLabel, Select, SelectChangeEvent} from "@mui/material";
import {Direction, Calculation, Unit, Value} from "../constants/AnalyticsDatesEnums.ts";

interface SelectFormControlProps {
    labelId: string,
    id: string,
    value: Direction | Unit | Calculation | Value,
    onChange: (event: SelectChangeEvent) => void,
    children: ReactNode
}

const SelectFormControl: FC<SelectFormControlProps> = ({labelId, id, value, onChange, children}) => (
    <FormControl variant="filled" sx={{m: 1, minWidth: 300}}>
        <InputLabel id={labelId}>Type</InputLabel>
        <Select labelId={labelId} id={id} value={value} onChange={onChange}>
            {children}
        </Select>
    </FormControl>
);

export default SelectFormControl;