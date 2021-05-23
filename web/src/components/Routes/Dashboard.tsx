import { useState } from "react";
import { useTitle } from "react-use";
import { useUserConfig } from "../../hooks/useUserConfig";
import { Menu } from "../Dashboard/Menu";
import { BttvForm } from "../Dashboard/RewardForms/BttvForm";
import { FfzForm } from "../Dashboard/RewardForms/FfzForm";

export function Dashboard() {
    useTitle("bitraft - Dashboard");
    const [renderKey, setRenderKey] = useState(1);
    const [userCfg, setUserConfig, fetchConfig] = useUserConfig(() => setRenderKey(renderKey + 1));

    if (!userCfg) {
        return null;
    }

    // force re-mount when switching the channel to manage, to re-render forms and their defaultValues
    return <div key={renderKey}>
        <Menu userConfig={userCfg} setUserConfig={setUserConfig} />
        <BttvForm userConfig={userCfg} setUserConfig={setUserConfig} fetchConfig={fetchConfig} />
        <FfzForm userConfig={userCfg} setUserConfig={setUserConfig} fetchConfig={fetchConfig} />
    </div>
}