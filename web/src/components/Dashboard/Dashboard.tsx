import { useTitle } from "react-use";
import { useUserConfig } from "../../hooks/useUserConfig";
import { Menu } from "./Menu";
import { PredictionLog } from "./PredictionLog";
import { BttvForm } from "./RewardForms/BttvForm";
import { Permissions } from "./Permissions";

export function Dashboard() {
    useTitle("bitraft - Dashboard");
    const [userCfg, setUserConfig] = useUserConfig();
    if (!userCfg) {
        return null;
    }

    return <div>
        <Menu userConfig={userCfg} setUserConfig={setUserConfig} />
        <div className="flex">
            <BttvForm userConfig={userCfg} />
            <PredictionLog />
            <Permissions userConfig={userCfg} />
        </div>
    </div>
}