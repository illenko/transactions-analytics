import {Route, Routes} from "react-router-dom";
import Layout from "../pages/Layout";
import Transactions from "../pages/Transactions";
import AnalyticsGroups from "../pages/AnalyticsGroups";
import AnalyticsDates from "../pages/AnalyticsDates";
import {FC} from "react";

const AppRoutes: FC = () => {
    return (
        <Routes>
            <Route path="/" element={<Layout/>}>
                <Route path="/transactions" element={<Transactions/>}/>
                <Route path="/analytics-groups" element={<AnalyticsGroups/>}/>
                <Route path="/analytics-dates" element={<AnalyticsDates/>}/>
            </Route>
        </Routes>
    );
};

export default AppRoutes;