package types

type Weapon struct {
	Name               string `json:"name"`
	Category           `json:"category"`
	PhysicalProp       `json:"physicalProp"`
	EngineProp         `json:"engineProp"`
	FuseAndWarheadProp `json:"fuseAndWarheadProp"`
	GuidanceProp       `json:"guidanceProp"`
	FlightProp         `json:"flightProp"`
}

type Category string

type PhysicalProp struct {
	Mass                     string `json:"mass,omitempty"`
	MassAtEndOfBoosterBurn   string `json:"massAtEndOfBoosterBurn,omitempty"`
	MassAtEndOfSustainerBurn string `json:"massAtEndOfSustainerBurn,omitempty"`
	Calibre                  string `json:"calibre,omitempty"`
	Length                   string `json:"length,omitempty"`
}

type EngineProp struct {
	ForceExertedByBooster      string `json:"forceExertedByBooster,omitempty"`
	BurnTimeOfBooster          string `json:"burnTimeOfBooster,omitempty"`
	RawAccelerationAtIgnition  string `json:"rawAccelerationAtIgnition,omitempty"`
	SpecificImpulseOfBooster   string `json:"specificImpulseOfBooster,omitempty"`
	DeltaSpeedOfBooster        string `json:"deltaSpeedOfBooster,omitempty"`
	BoosterStartDelay          string `json:"boosterStartDelay,omitempty"`
	ForceExertedBySustainer    string `json:"forceExertedBySustainer,omitempty"`
	BurnTimeOfSustainer        string `json:"burnTimeOfSustainer,omitempty"`
	SpecificImpulseOfSustainer string `json:"specificImpulseOfSustainer,omitempty"`
	DeltaSpeedOfSustainer      string `json:"deltaSpeedOfSustainer,omitempty"`
	TotalDeltaSpeed            string `json:"totalDeltaSpeed,omitempty"`
}

type FuseAndWarheadProp struct {
	ExplosiveMass                string `json:"explosiveMass,omitempty"`
	TandemCharge                 string `json:"tandemCharge,omitempty"`
	Penetration                  string `json:"penetration,omitempty"`
	ProximityFuse                string `json:"proximityFuse,omitempty"`
	ProximityFuseRange           string `json:"proximityFuseRange,omitempty"`
	ProximityFuseArmingDistance  string `json:"proximityFuseArmingDistance,omitempty"`
	ProximityFuseShellDetection  string `json:"proximityFuseShellDetection,omitempty"`
	ProximityFuseMinimumAltitude string `json:"proximityFuseMinimumAltitude,omitempty"`
	ProximityFuseDelay           string `json:"proximityFuseDelay,omitempty"`
}

type GuidanceProp struct {
	Zoom                                          string `json:"zoom,omitempty"`
	GuidanceType                                  string `json:"guidanceType,omitempty"`
	GuidanceStartDelay                            string `json:"guidanceStartDelay,omitempty"`
	GuidanceDuration                              string `json:"guidanceDuration,omitempty"`
	GuidanceRange                                 string `json:"guidanceRange,omitempty"`
	LaunchSector                                  string `json:"launchSector,omitempty"`
	ControlConeFOV                                string `json:"controlConeFOV,omitempty"`
	AimTrackingSensitivity                        string `json:"aimTrackingSensitivity,omitempty"`
	MaximumAngleAllowedBetweenMissileAndCrosshair string `json:"maximumAngleAllowedBetweenMissileAndCrosshair,omitempty"`
	SeekerWarmUpTime                              string `json:"seekerWarmUpTime,omitempty"`
	SeekerSearchDuration                          string `json:"seekerSearchDuration,omitempty"`
	FieldOfView                                   string `json:"fieldOfView,omitempty"`
	OpticSightFieldOfView                         string `json:"opticSightFieldOfView,omitempty"`
	GimbalLimit                                   string `json:"gimbalLimit,omitempty"`
	TrackRate                                     string `json:"trackRate,omitempty"`
	UncageSeekerBeforeLaunch                      string `json:"uncageSeekerBeforeLaunch,omitempty"`
	MaximumLockAngleBeforeLaunch                  string `json:"maximumLockAngleBeforeLaunch,omitempty"`
	MinimumAngleBetweenSeekerAndSunForNotCapture  string `json:"minimumAngleBetweenSeekerAndSunForNotCapture,omitempty"`
	CanLockGround                                 string `json:"canLockGround,omitempty"`
	LockOnRangeGround                             string `json:"lockOnRangeGround,omitempty"`
	LockOnRangeVehicle                            string `json:"lockOnRangeVehicle,omitempty"`
	LockOnRangeFromRearAspect                     string `json:"lockOnRangeFromRearAspect,omitempty"`
	FlareDetectionRange                           string `json:"flareDetectionRange,omitempty"`
	IRCMDetectionRange                            string `json:"IRCMDetectionRange,omitempty"`
	DIRCMDetectionRange                           string `json:"DIRCMDetectionRange,omitempty"`
	HeadOnLockOnRangeAgainstAfterburnerTarget     string `json:"headOnLockOnRangeAgainstAfterburnerTarget,omitempty"`
	IRCCM                                         string `json:"IRCCM,omitempty"`
	IRCCMType                                     string `json:"IRCCMType,omitempty"`
	IRCCMFieldOfView                              string `json:"IRCCMFieldOfView,omitempty"`
	IRCCMRejectionThreshold                       string `json:"IRCCMRejectionThreshold,omitempty"`
	IRCCMReactionTime                             string `json:"IRCCMReactionTime,omitempty"`
	LockOnRangeFromAllAspect                      string `json:"lockOnRangeFromAllAspect,omitempty"`
	CountermeasureDetectionRange                  string `json:"countermeasureDetectionRange,omitempty"`
	MaximumBreakLockTime                          string `json:"maximumBreakLockTime,omitempty"`
	CanBeSlavedToRadar                            string `json:"canBeSlavedToRadar,omitempty"`
	CanLockAfterLaunch                            string `json:"canLockAfterLaunch,omitempty"`
	Band                                          string `json:"band,omitempty"`
	AngularSpeedRejectionThreshold                string `json:"angularSpeedRejectionThreshold,omitempty"`
	AngularRejectionThresholdRange                string `json:"angularRejectionThresholdRange,omitempty"`
	AccelerationRejectionThresholdRange           string `json:"accelerationRejectionThresholdRange,omitempty"`
	SidelobeAttenuation                           string `json:"sidelobeAttenuation,omitempty"`
	TransmitterPower                              string `json:"transmitterPower,omitempty"`
	TransmitterAngleOfHalfSensitivity             string `json:"transmitterAngleOfHalfSensitivity,omitempty"`
	TransmitterSidelobeSensitivity                string `json:"transmitterSidelobeSensitivity,omitempty"`
	ReceiverAngleOfHalfSensitivity                string `json:"receiverAngleOfHalfSensitivity,omitempty"`
	ReceiverSidelobeSensitivity                   string `json:"receiverSidelobeSensitivity,omitempty"`
	DistanceMinimumValue                          string `json:"distanceMinimumValue,omitempty"`
	DistanceMaximumValue                          string `json:"distanceMaximumValue,omitempty"`
	DistanceWidth                                 string `json:"distanceWidth,omitempty"`
	DistanceMinimumSignalGate                     string `json:"distanceMinimumSignalGate,omitempty"`
	DistanceRefWidth                              string `json:"distanceRefWidth,omitempty"`
	DistanceGateSearchRange                       string `json:"distanceGateSearchRange,omitempty"`
	DistanceGateAlphaFilter                       string `json:"distanceGateAlphaFilter,omitempty"`
	DistanceGateBetaFilter                        string `json:"distanceGateBetaFilter,omitempty"`
	DopplerSpeedMinimumValue                      string `json:"dopplerSpeedMinimumValue,omitempty"`
	DopplerSpeedMaximumValue                      string `json:"dopplerSpeedMaximumValue,omitempty"`
	DopplerSpeedWidth                             string `json:"dopplerSpeedWidth,omitempty"`
	DopplerSpeedRefWidth                          string `json:"dopplerSpeedRefWidth,omitempty"`
	DopplerSpeedMinimumSignalGate                 string `json:"dopplerSpeedMinimumSignalGate,omitempty"`
	DopplerSpeedGateSearchRange                   string `json:"dopplerSpeedGateSearchRange,omitempty"`
	DopplerSpeedGateAlphaFilter                   string `json:"dopplerSpeedGateAlphaFilter,omitempty"`
	DopplerSpeedGateBetaFilter                    string `json:"dopplerSpeedGateBetaFilter,omitempty"`
	ProportionalNavigationMultiplier              string `json:"proportionalNavigationMultiplier,omitempty"`
	BaseIndicatedAirSpeed                         string `json:"baseIndicatedAirSpeed,omitempty"`
	PIDProportionalTerm                           string `json:"PIDProportionalTerm,omitempty"`
	PIDIntegralTerm                               string `json:"PIDIntegralTerm,omitempty"`
	PIDIntegralTermLimit                          string `json:"PIDIntegralTermLimit,omitempty"`
	PIDDerivativeTerm                             string `json:"PIDDerivativeTerm,omitempty"`
	InertialGuidanceDriftSpeed                    string `json:"inertialGuidanceDriftSpeed,omitempty"`
	InertialNavigation                            string `json:"inertialNavigation,omitempty"`
	DistanceGate                                  string `json:"distanceGate,omitempty"`
	InertialNavigationDriftSpeed                  string `json:"inertialNavigationDriftSpeed,omitempty"`
}

type FlightProp struct {
	MaximumLaunchAngleHorizontalVertical               string `json:"maximumLaunchAngleHorizontalVertical,omitempty"`
	AimSensitivity                                     string `json:"aimSensitivity,omitempty"`
	MaximumAxisValues                                  string `json:"maximumAxisValues,omitempty"`
	MaximumFinAngleOfAttack                            string `json:"maximumFinAngleOfAttack,omitempty"`
	FinsLateralAcceleration                            string `json:"finsLateralAcceleration,omitempty"`
	MaximumAOA                                         string `json:"maximumAOA,omitempty"`
	MaximumFinLateralAcceleration                      string `json:"maximumFinLateralAcceleration,omitempty"`
	WingAreaMultiplier                                 string `json:"wingAreaMultiplier,omitempty"`
	MaximumLateralAcceleration                         string `json:"maximumLateralAcceleration,omitempty"`
	StartSpeed                                         string `json:"startSpeed,omitempty"`
	MaximumSpeed                                       string `json:"maximumSpeed,omitempty"`
	MinimumRange                                       string `json:"minimumRange,omitempty"`
	MaximumFlightRange                                 string `json:"maximumFlightRange,omitempty"`
	Tracer                                             string `json:"tracer,omitempty"`
	LoadFactorLimitAtLaunch                            string `json:"loadFactorLimitAtLaunch,omitempty"`
	MaximumOverLoad                                    string `json:"maximumOverLoad,omitempty"`
	SeaSkimming                                        string `json:"seaSkimming,omitempty"`
	FlightTimeUntilGuidanceStarts                      string `json:"flightTimeUntilGuidanceStarts,omitempty"`
	FlightTimeWhenPullLimit30                          string `json:"flightTimeWhenPullLimit30%,omitempty"`
	FlightTimeWhenPullLimit40                          string `json:"flightTimeWhenPullLimit40%,omitempty"`
	FlightTimeWhenPullLimit100                         string `json:"flightTimeWhenPullLimit100%,omitempty"`
	Loft                                               string `json:"loft,omitempty"`
	LoftAngle                                          string `json:"loftAngle,omitempty"`
	TargetElevation                                    string `json:"targetElevation,omitempty"`
	MaximumTargetAngularChange                         string `json:"maximumTargetAngularChange,omitempty"`
	ThrustVectoring                                    string `json:"thrustVectoring,omitempty"`
	ThrustVectoringAngle                               string `json:"thrustVectoringAngle,omitempty"`
	StartingGLimit                                     string `json:"startingGLimit,omitempty"`
	ETAToImpactWhenPropMultiplierReachesXPercentage30  string `json:"ETAToImpactWhenPropMultiplierReachesXPercentage30%,omitempty"`
	ETAToImpactWhenPropMultiplierReachesXPercentage50  string `json:"ETAToImpactWhenPropMultiplierReachesXPercentage50%,omitempty"`
	ETAToImpactWhenPropMultiplierReachesXPercentage80  string `json:"ETAToImpactWhenPropMultiplierReachesXPercentage80%,omitempty"`
	ETAToImpactWhenPropMultiplierReachesXPercentage90  string `json:"ETAToImpactWhenPropMultiplierReachesXPercentage90%,omitempty"`
	ETAToImpactWhenPropMultiplierReachesXPercentage100 string `json:"ETAToImpactWhenPropMultiplierReachesXPercentage100%,omitempty"`
	ETAToImpactWhenPropMultiplierReachesXPercentageX   string `json:"ETAToImpactWhenPropMultiplierReachesXPercentageX%,omitempty"`
}

func NewWeapon(weapon *Weapon) *Weapon {
	return &Weapon{
		Name:     weapon.Name,
		Category: weapon.Category,
		PhysicalProp: PhysicalProp{
			Mass: weapon.Mass,
		},
	}
}
