

/**
 * RDPClient is a minimal RDP client for nuclei scripts.
 * @example
 * ```javascript
 * const rdp = require('nuclei/rdp');
 * const client = new rdp.Client();
 * ```
 */
export class RDPClient {
    

    // Constructor of RDPClient
    constructor() {}
    /**
    * IsRDP checks if the given host and port are running rdp server.
    * If connection is successful, it returns true.
    * If connection is unsuccessful, it returns false and error.
    * The Name of the OS is also returned if the connection is successful.
    * @example
    * ```javascript
    * const rdp = require('nuclei/rdp');
    * const isRDP = rdp.IsRDP('acme.com', 3389);
    * log(toJSON(isRDP));
    * ```
    */
    public IsRDP(host: string, port: number): IsRDPResponse | null {
        return null;
    }
    

    /**
    * CheckRDPAuth checks if the given host and port are running rdp server
    * with authentication and returns their metadata.
    * If connection is successful, it returns true.
    * @example
    * ```javascript
    * const rdp = require('nuclei/rdp');
    * const checkRDPAuth = rdp.CheckRDPAuth('acme.com', 3389);
    * log(toJSON(checkRDPAuth));
    * ```
    */
    public CheckRDPAuth(host: string, port: number): CheckRDPAuthResponse | null {
        return null;
    }
    

}



/**
 * CheckRDPAuthResponse is the response from the CheckRDPAuth function.
 * this is returned by CheckRDPAuth function.
 * @example
 * ```javascript
 * const rdp = require('nuclei/rdp');
 * const checkRDPAuth = rdp.CheckRDPAuth('acme.com', 3389);
 * log(toJSON(checkRDPAuth));
 * ```
 */
export interface CheckRDPAuthResponse {
    
    PluginInfo?: ServiceRDP,
    
    Auth?: boolean,
}



/**
 * IsRDPResponse is the response from the IsRDP function.
 * this is returned by IsRDP function.
 * @example
 * ```javascript
 * const rdp = require('nuclei/rdp');
 * const isRDP = rdp.IsRDP('acme.com', 3389);
 * log(toJSON(isRDP));
 * ```
 */
export interface IsRDPResponse {
    
    IsRDP?: boolean,
    
    OS?: string,
}



/**
 * ServiceRDP Interface
 */
export interface ServiceRDP {
    
    NetBIOSComputerName?: string,
    
    NetBIOSDomainName?: string,
    
    DNSComputerName?: string,
    
    DNSDomainName?: string,
    
    ForestName?: string,
    
    OSFingerprint?: string,
    
    OSVersion?: string,
    
    TargetName?: string,
}

