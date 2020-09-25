package t

type Key string

const (
	NO         Key = "No"
	YES        Key = "Yes"
	CANCEL     Key = "Cancel"
	CANCELED   Key = "Canceled"
	COMPLETED  Key = "Completed"
	CONFIRM    Key = "Confirm"
	PAYAMOUNT  Key = "PayAmount"
	FAILURE    Key = "Failure"
	PROCESSING Key = "Processing"
	WITHDRAW   Key = "Withdraw"
	ERROR      Key = "Error"
	CHECKING   Key = "Checking"
	TXPENDING  Key = "TxPending"
	TXCANCELED Key = "TxCanceled"
	UNEXPECTED Key = "Unexpected"

	CLAIMFAILED Key = "ClaimFailed"

	CALLBACKCOINFLIPWINNER Key = "CallbackCoinflipWinner"
	CALLBACKWINNER         Key = "CallbackWinner"
	CALLBACKERROR          Key = "CallbackError"
	CALLBACKEXPIRED        Key = "CallbackExpired"
	CALLBACKATTEMPT        Key = "CallbackAttempt"
	CALLBACKSENDING        Key = "CallbackSending"

	INLINEINVOICERESULT  Key = "InlineInvoiceResult"
	INLINEGIVEAWAYRESULT Key = "InlineGiveawayResult"
	INLINEGIVEFLIPRESULT Key = "InlineGiveflipResult"
	INLINECOINFLIPRESULT Key = "InlineCoinflipResult"
	INLINEHIDDENRESULT   Key = "InlineHiddenResult"

	LNURLUNSUPPORTED      Key = "LnurlUnsupported"
	LNURLERROR            Key = "LnurlError"
	LNURLAUTHSUCCESS      Key = "LnurlAuthSuccess"
	LNURLPAYPROMPT        Key = "LnurlPayPrompt"
	LNURLPAYPROMPTCOMMENT Key = "LnurlPayPromptComment"
	LNURLPAYSUCCESS       Key = "LnurlPaySuccess"
	LNURLPAYMETADATA      Key = "LnurlPayMetadata"
	LNURLPAYCOMMENT       Key = "LnurlPayComment"

	USERALLOWED       Key = "UserAllowed"
	SPAMFILTERMESSAGE Key = "SpamFilterMessage"

	RENAMABLEMSG      Key = "RenamableMsg"
	RENAMEPROMPT      Key = "RenamePrompt"
	GROUPNOTRENAMABLE Key = "GroupNotRenamable"

	PAYMENTFAILED       Key = "PaymentFailed"
	PAIDMESSAGE         Key = "PaidMessage"
	DBERROR             Key = "DBError"
	INSUFFICIENTBALANCE Key = "InsufficientBalance"
	RATELIMIT           Key = "RateLimit"
	OVERQUOTA           Key = "OverQuota"

	PAYMENTRECEIVED      Key = "PaymentReceived"
	FAILEDTOSAVERECEIVED Key = "FailedToSaveReceived"

	SPAMMYMSG           Key = "SpammyMsg"
	COINFLIPSENABLEDMSG Key = "CoinflipsEnabledMsg"
	LANGUAGEMSG         Key = "LanguageMsg"
	TICKETMSG           Key = "TicketMsg"
	FREEJOIN            Key = "FreeJoin"

	APPBALANCE Key = "AppBalance"

	HELPINTRO   Key = "HelpIntro"
	HELPSIMILAR Key = "HelpSimilar"
	HELPMETHOD  Key = "HelpMethod"

	RECEIVEHELP Key = "receiveHelp"

	PAYHELP Key = "payHelp"

	SENDHELP Key = "sendHelp"

	TRANSACTIONSHELP Key = "transactionsHelp"

	BALANCEHELP Key = "balanceHelp"

	GIVEAWAYHELP            Key = "giveawayHelp"
	GIVEAWAYMSG             Key = "GiveAwayMsg"
	GIVEAWAYCLAIM           Key = "GiveAwayClaim"
	GIVEAWAYSATSGIVENPUBLIC Key = "GiveawaySatsGivenPublic"

	COINFLIPHELP      Key = "coinflipHelp"
	COINFLIPWINNERMSG Key = "CoinflipWinnerMsg"
	COINFLIPGIVERMSG  Key = "CoinflipGiverMsg"
	COINFLIPAD        Key = "CoinflipAd"
	COINFLIPJOIN      Key = "CoinflipJoin"

	GIVEFLIPHELP      Key = "giveflipHelp"
	GIVEFLIPMSG       Key = "GiveFlipMsg"
	GIVEFLIPWINNERMSG Key = "GiveflipWinnerMsg"
	GIVEFLIPAD        Key = "GiveflipAd"
	GIVEFLIPJOIN      Key = "GiveflipJoin"

	FUNDRAISEHELP        Key = "fundraiseHelp"
	FUNDRAISEAD          Key = "FundraiseAd"
	FUNDRAISEJOIN        Key = "FundraiseJoin"
	FUNDRAISECOMPLETE    Key = "FundraiseComplete"
	FUNDRAISERECEIVERMSG Key = "FundraiseReceiverMsg"
	FUNDRAISEGIVERMSG    Key = "FundraiseGiverMsg"

	LIGHTNINGATMHELP       Key = "lightningatmHelp"
	BLUEWALLETHELP         Key = "bluewalletHelp"
	APIPASSWORDUPDATEERROR Key = "APIPasswordUpdateError"
	APICREDENTIALS         Key = "APICredentials"

	HIDEHELP             Key = "hideHelp"
	REVEALHELP           Key = "revealHelp"
	HIDDENREVEALBUTTON   Key = "HiddenRevealButton"
	HIDDENDEFAULTPREVIEW Key = "HiddenDefaultPreview"
	HIDDENWITHID         Key = "HiddenWithId"
	HIDDENSOURCEMSG      Key = "HiddenSourceMsg"
	HIDDENREVEALMSG      Key = "HiddenRevealMsg"
	HIDDENMSGNOTFOUND    Key = "HiddenMsgNotFound"
	HIDDENSHAREBTN       Key = "HiddenShareBtn"

	ETLENEUMHELP          Key = "etleneumHelp"
	ETLENEUMACCOUNT       Key = "EtleneumAccount"
	ETLENEUMHISTORY       Key = "EtleneumHistory"
	ETLENEUMCONTRACT      Key = "EtleneumContract"
	ETLENEUMCONTRACTSTATE Key = "EtleneumContractState"
	ETLENEUMCALL          Key = "EtleneumCall"
	ETLENEUMCONTRACTS     Key = "EtleneumContracts"
	ETLENEUMSUBSCRIBED    Key = "EtleneumSubscribed"
	ETLENEUMCONTRACTEVENT Key = "EtleneumContractEvent"

	BITREFILLHELP            Key = "bitrefillHelp"
	BITREFILLINVENTORYHEADER Key = "BitrefillInventoryHeader"
	BITREFILLPACKAGESHEADER  Key = "BitrefillPackagelHeader"
	BITREFILLNOPROVIDERS     Key = "BitrefillNoProviders"
	BITREFILLCONFIRMATION    Key = "BitrefillConfirmation"
	BITREFILLFAILEDSAVE      Key = "BitrefillFailedToSave"
	BITREFILLPURCHASEDONE    Key = "BitrefillPurchaseDone"
	BITREFILLPURCHASEFAILED  Key = "BitrefillPurchaseFailed"
	BITREFILLCOUNTRYSET      Key = "BitrefillCountrySet"
	BITREFILLINVALIDCOUNTRY  Key = "BitrefillInvalidCountry"

	SATELLITEHELP Key = "satelliteHelp"
	SATELLITEPAID Key = "SatellitePaid"

	FUNDBTCHELP   Key = "fundbtcHelp"
	FUNDBTCFAIL   Key = "fundbtcFail"
	FUNDBTCFINISH Key = "fundbtcFinish"

	BITCLOUDSHELP           Key = "bitcloudsHelp"
	BITCLOUDSCREATEHEADER   Key = "BitcloudsCreateHeader"
	BITCLOUDSCREATED        Key = "BitcloudsCreated"
	BITCLOUDSSTOPPEDWAITING Key = "BitcloudsStoppedWaiting"
	BITCLOUDSNOHOSTS        Key = "BitcloudsNoHosts"
	BITCLOUDSHOSTSHEADER    Key = "BitcloudsHostsHeader"
	BITCLOUDSSTATUS         Key = "BitcloudsStatus"
	BITCLOUDSREMINDER       Key = "BitcloudsReminder"

	SKYPEHELP Key = "skypeHelp"
	RUBHELP   Key = "rubHelp"

	GIFTSHELP       Key = "giftsHelp"
	GIFTSCREATED    Key = "GiftsCreated"
	GIFTSFAILEDSAVE Key = "GiftsFailedSave"
	GIFTSLIST       Key = "GiftsList"
	GIFTSSPENTEVENT Key = "GiftsSpentEvent"

	SATS4ADSHELP       Key = "sats4adsHelp"
	SATS4ADSTOGGLE     Key = "Sats4adsToggle"
	SATS4ADSBROADCAST  Key = "Sats4adsBroadcast"
	SATS4ADSSTART      Key = "Sats4adsStart"
	SATS4ADSPRICETABLE Key = "Sats4adsPriceTable"
	SATS4ADSADFOOTER   Key = "Sats4adsAdFooter"
	SATS4ADSVIEWED     Key = "Viewed"

	ETLENEUMFAILEDTOPAY Key = "EtleneumFailedToPay"

	TOGGLEHELP Key = "toggleHelp"

	HELPHELP Key = "helpHelp"

	STOPHELP Key = "stopHelp"

	PAYPROMPT          Key = "PayPrompt"
	FAILEDDECODE       Key = "FailedDecode"
	BALANCEMSG         Key = "BalanceMsg"
	TAGGEDBALANCEMSG   Key = "TaggedBalanceMsg"
	FAILEDUSER         Key = "FailedUser"
	LOTTERYMSG         Key = "LotteryMsg"
	INVALIDPARTNUMBER  Key = "InvalidPartNumber"
	INVALIDAMOUNT      Key = "InvalidAmount"
	USERSENTTOUSER     Key = "UserSentToUser"
	USERSENTYOUSATS    Key = "UserSentYouSats"
	RECEIVEDSATSANON   Key = "ReceivedSatsAnon"
	FAILEDSEND         Key = "FailedSend"
	QRCODEFAIL         Key = "QRCodeFail"
	SAVERECEIVERFAIL   Key = "SaveReceiverFail"
	CANTSENDNORECEIVER Key = "CantSendNoReceiver"
	GIVERCANTJOIN      Key = "GiverCantJoin"
	CANTJOINTWICE      Key = "CantJoinTwice"
	CANTCANCEL         Key = "CantCancel"
	FAILEDINVOICE      Key = "FailedInvoice"
	INVALIDAMT         Key = "InvalidAmt"
	STOPNOTIFY         Key = "StopNotify"
	WELCOME            Key = "Welcome"
	WRONGCOMMAND       Key = "WrongCommand"
	RETRACTQUESTION    Key = "RetractQuestion"
	RECHECKPENDING     Key = "RecheckPending"

	TXNOTFOUND Key = "TxNotFound"
	TXINFO     Key = "TxInfo"
	TXLIST     Key = "TxList"
	TXLOG      Key = "TxLog"

	TUTORIALWALLET  Key = "TutorialWallet"
	TUTORIALBLUE    Key = "TutorialBlue"
	TUTORIALAPPS    Key = "TutorialApps"
	TUTORIALTWITTER Key = "TutorialTwitter"
)
