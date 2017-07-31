pragma solidity ^0.4.0;

import "./abstracts/CourtroomAbstract.sol";
import "./sampletoken.sol";
import "./abstracts/trialrulesabstract.sol";


contract SwearGame is SwearGameAbstract {

    uint  public playerCount;
    SampleToken public token;
    TrialRulesAbstract public trialRules;

    struct Case {
        address plaintiff;
        bytes32 serviceId;
        uint8 status;
        uint8 valid;
    }

    struct Deposit {
        bool inDepositPeriod;
        uint vestingPeriod;
        uint depositedAmount;
    }
    mapping(address=>Deposit) public deposits;
    mapping(address => bool) public players;
    //id map to Case
    mapping(bytes32 => Case)  OpenCases;
    mapping(address => uint256)  OpenCasesNumber;
    mapping(address => bytes32[]) public ids;

    /// @notice SwearGame - Swear game constructor this function is called along with
    /// the contract deployment time.
    /// @param _token - address of the token contract
    /// @param _trialRules - address of the trial specific rules contract
    /// @return WitnessAbstract - return a witness contract instance
    function SwearGame(address _token,address _trialRules) {
        token = SampleToken(_token);
        trialRules = TrialRulesAbstract(_trialRules);
        playerCount = 0;
    }

    /// @notice _newCase - open a new case and add it to OpenCases
    ///
    /// @param _plaintiff  - the plaintiff address for the case
    /// @param _serviceId - service id related to the case
    /// @return _status - the status of the case
    function _newCase(address _plaintiff,bytes32 _serviceId,uint8 _status) private returns (bytes32 id) {
        id = sha3(_plaintiff,_serviceId, now);
        if (OpenCases[id].valid != 0)
           return 0x0;
        OpenCases[id] = Case(
            _plaintiff,
            _serviceId,
            _status,
            1
        );
        return id;
    }

    function isValid(bytes32 id) private constant returns (bool) {
        return (OpenCases[id].valid != 0);
    }

    function setStatus(bytes32 id,uint8 status) private constant  returns (bool) {
        if (msg.sender == owner)
            throw;
        OpenCases[id].status = status;
        return true;
    }

    function resolveCase(bytes32 _id) private {

        if (OpenCases[_id].status == uint8(TrialRulesAbstract.Status.UNCHALLENGED))
            throw;
        OpenCases[_id].plaintiff = 0;
        OpenCases[_id].valid = 0;
        OpenCases[_id].status = uint8(TrialRulesAbstract.Status.UNCHALLENGED);
    }

    function getCase(bytes32 _id) private returns (address plaintiff,bytes32 serviceId) {
        return (OpenCases[_id].plaintiff,OpenCases[_id].serviceId);
    }

    function verdict(bytes32 _id,uint8 status,address plaintiff) private returns(bool) {

        if (status == uint8(TrialRulesAbstract.Status.NOT_GUILTY)) {
            CaseResolved(
                _id,
                plaintiff,
                0,
                status
            );
            return false;
        }
        uint reward = trialRules.getReward();
        bool caseCompensated = compensate(plaintiff,reward);
        resolveCase(_id);
        unRegister(plaintiff);
        CaseResolved(
            _id,
            plaintiff,
            reward,
            status
        );
        return caseCompensated;
    }

    /// @notice leaveGame - dismiss a player from the game (unregister)
    /// allow only plaintiff which do not have openCases on it name to leave game
    /// @param _player  - the player address
    function leaveGame(address _player) {

        for (uint256 i = 0;i<ids[msg.sender].length;i++) {
        //allow only plaintiff which do not have openCases on it name to leave game
            require(OpenCases[ids[msg.sender][i]].valid == 0);
        }
        return unRegister(msg.sender);
    }

    /// @notice getStatus - return the trial status of a case
    ///
    /// @param id  - case id
    /// @return  status  - the status of a case
    function getStatus(bytes32 id) public constant returns (uint8 status) {
        return OpenCases[id].status;
    }

    /// @notice newCase - open a new case for a service id
    ///
    /// the function require that the msg sender is already register to the game.
    /// @param serviceId  - service id
    /// @return bool - true for successful operation.
    function newCase(bytes32 serviceId) public returns (bool) {

        require(players[msg.sender]);
        bytes32 id = _newCase(msg.sender,serviceId,uint8(trialRules.getInitialStatus()));
        if (id == 0x0)
            return false;
        OpenCasesNumber[msg.sender]++;
        OpenCasesNumber[owner]++;
        ids[msg.sender].push(id);

        return true;
    }

    /// @notice trial - initiate or restart a trial proccess for a certain case
    ///
    /// the function requiere that the case is a valid one.
    /// @param id  - case id
    /// @return bool - true for successful operation.
    function trial(bytes32 id) public returns (bool) {

        require(players[msg.sender]);
        require(isValid(id));
        _trial(id);
        return true;
    }

    function proceed() private returns (WitnessAbstract.Status) {
        return WitnessAbstract.Status.PENDING;
    }

    function _trial(bytes32 id) private {

        uint8 status = getStatus(id);
        var(plaintiff,serviceId) = getCase(id);
        while (status != uint8(TrialRulesAbstract.Status.UNCHALLENGED)) {
            WitnessAbstract witness = trialRules.getWitness(status);
            WitnessAbstract.Status outcome;
            if (witness == WitnessAbstract(0x0)) {
                outcome = proceed();
                return;
                } else {
                bool expired = trialRules.expired(id,status);
                if (witness.isEvidenceSubmitted(id,serviceId,plaintiff) && !expired) {
                    outcome = witness.testimonyFor(id,serviceId,plaintiff);
                    }else {
                    if (trialRules.startGracePeriod(id,status)||(!expired)) {
                        outcome = WitnessAbstract.Status.PENDING;
                        }else {
                        outcome = WitnessAbstract.Status.INVALID;
                        }
                    }
                }
            if (outcome == WitnessAbstract.Status.PENDING) {
                return;
                }
            status = trialRules.getStatus(uint8(outcome),status);
            setStatus(id,status);
            if ((status == uint8(TrialRulesAbstract.Status.GUILTY))||
                (status == uint8(TrialRulesAbstract.Status.NOT_GUILTY))){
                verdict(id,status,plaintiff);
                status = uint8(TrialRulesAbstract.Status.UNCHALLENGED);
                OpenCasesNumber[plaintiff]--;
                OpenCasesNumber[owner]--;
                setStatus(id,status);
                }
        }
    }

    function unRegister(address _player) {

        require(players[_player]);
        PlayerLeftGame(_player);
        players[_player] = false;
        playerCount--;
    }

    /// @notice register - register a player to the game
    ///
    /// The function will throw if the player is already register or there is not
    /// enough deposit in the contract to ensure the player could be compensated for the
    /// case of a valid case.
    /// @param _player  - the player address
    /// @return bool registered - true for success registration.
    function register(address _player) onlyOwner public returns (bool) {

        require(!players[_player]);
        uint reward = trialRules.getReward();
        if (playerCount == 0) {
            require(deposits[owner].depositedAmount >= reward);
        }else if ((deposits[owner].depositedAmount / playerCount) < reward) {
            AdditionalDepositRequired(deposits[owner].depositedAmount);
            throw;
        }
        players[_player] = true;
        playerCount++;
        NewPlayer(_player);
        return true;
    }

    function deposit(uint epochs) payable returns (bool) {

        require(token.transferFrom(msg.sender, address(this), msg.value));
        if (deposits[msg.sender].inDepositPeriod) {
            deposits[msg.sender].depositedAmount += msg.value;
        }else {
            deposits[msg.sender] = Deposit({inDepositPeriod: true, vestingPeriod: block.number + epochs * trialRules.getEpoch(), depositedAmount: msg.value});
        }
        DepositStaked(msg.value, deposits[msg.sender].depositedAmount);
        return true;
    }

    function collectDeposit() external returns (bool) {

        require(OpenCasesNumber[msg.sender] == 0);//check if there is no open case for the specific caller.
        Deposit storage depositInfo = deposits[msg.sender];
        if (depositInfo.inDepositPeriod && depositInfo.vestingPeriod <= block.number) {
            uint toTransfer = depositInfo.depositedAmount;
            deposits[msg.sender] = Deposit(false, 0, 0);
            token.transferFrom(address(this),msg.sender, toTransfer);
            return true;
          }
        return false;
    }

    function isRegister(address player) returns (bool) {
        return players[msg.sender];
    }

    function compensate(address _beneficiary,uint reward) private returns(bool compensated) {

        compensated = token.transferFrom(address(this), _beneficiary, reward);
        require(compensated);
        deposits[owner].depositedAmount -= reward;
        Compensate(_beneficiary,reward);
        return compensated;
    }

    event Decision(string decide);
    event NewCaseOpened(bytes32 id, address plaintiff);
    event NewEvidenceSubmitted(bytes32 id, address plaintiff);
    event CaseResolved(bytes32 id, address plaintiff, uint reward,uint8 status);
    event Payment(address from,address to ,uint256 value);
    event NewPlayer(address playerId);
    event AdditionalDepositRequired(uint256 deposit);
    event DepositStaked(uint depositAmount, uint deposit);
    event Compensate(address recipient, uint reward);
    event PlayerLeftGame(address playerId);

}
