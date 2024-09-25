
package wirepod_ttr

import (
    "errors"
    "fmt"

    "github.com/fforchino/vector-go-sdk/pkg/vector"
    "github.com/kercre123/wire-pod/chipper/pkg/logger"
    "github.com/kercre123/wire-pod/chipper/pkg/vars"
)

// FindRobotByESN retrieves a robot instance by its ESN.
func FindRobotByESN(esn string) (*vector.Vector, error) {
    if esn == "" {
        return nil, errors.New("ESN cannot be empty")
    }

    for _, r := range vars.BotInfo.Robots {
        if r.Esn == esn {
            guid := r.GUID
            target := fmt.Sprintf("%s:443", r.IPAddress)

            robot, err := vector.New(vector.WithSerialNo(esn), vector.WithToken(guid), vector.WithTarget(target))
            if err != nil {
                logger.Println("Failed to create robot with ESN: %s, GUID: %s, Target: %s. Error: %v", esn, guid, target, err)
                return nil, fmt.Errorf("failed to create robot: %w", err)
            }
            return robot, nil
        }
    }

    logger.Println("No robot found with ESN:", esn)
    return nil, errors.New("no robot found with the given ESN")
}

//
// get_robot_id.go - END
//
