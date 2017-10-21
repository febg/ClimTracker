//
//  SignInViewController.swift
//  ClimbTracker
//
//  Created by Felipe Ballesteros on 2017-10-10.
//  Copyright © 2017 Felipe Ballesteros. All rights reserved.
//

import UIKit
import SwiftSpinner

class SignInViewController: UIViewController {

    var uemail = ""
    var uStatus = ""
    var upassword = ""
    var uName = ""
    var uID = ""
    var uStatusFlag = false
    var timer:                  Timer?
    var connectionAttempts = 0
    
    @IBOutlet weak var email: UITextField!
    
    @IBOutlet weak var password: UITextField!
    
    @IBAction func signInB(_ sender: Any) {
        if (email.text! != "" && password.text! != "") {
            SwiftSpinner.show(delay: 0, title: "Connecting to Server", animated: true)
          
            uemail = email.text!
            upassword = password.text!
            postRequest("http://standard-lb-1065564336.us-east-2.elb.amazonaws.com/logIn/")
            email.text = ""
            password.text = ""
            runTimer()
        }
        
        //Check for user in data base and grant/deny access
        
    }
    func saveDefaults(){
        UserDefaults.standard.setValue(uemail, forKey: "email")
        UserDefaults.standard.setValue(upassword, forKey: "password")
        UserDefaults.standard.setValue(uName, forKey: "name")
        UserDefaults.standard.setValue(uID, forKey: "tableID")
        
    }
    override func viewDidLoad() {
        super.viewDidLoad()
        self.view.addGestureRecognizer(UITapGestureRecognizer(target: self.view, action: #selector(UIView.endEditing(_:))))
        // Do any additional setup after loading the view.
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }
    
    func postRequest(_ url:String)
    {
        let url:NSURL = NSURL(string: url)!
        let session = URLSession.shared
        
        let request = NSMutableURLRequest(url: url as URL)
        request.httpMethod = "POST"
        
        let params = "email="+uemail+"&password="+upassword
        //let pass = "pass=test"
        request.httpBody = params.data(using: String.Encoding.utf8)
        //request.httpBody = pass.data(using: String.Encoding.utf8)
        
        let task = session.dataTask(with: request as URLRequest) {
            (
            data, response, error) in
            do{
            
            
            if let data = data, let json = try JSONSerialization.jsonObject(with: data) as? [String: Any]{
                if let status = json["Status"] as! String?{
                    print(json)
                    self.uStatus = status
                    self.uStatusFlag = true
                }
                
            }
            
            } catch {
                print("Error deserializing JSON: \(error)")
                
            }
        }
        
        task.resume()
        
    }
    
    func runTimer() {
        timer = Timer.scheduledTimer(timeInterval: 1, target:self, selector: #selector(updateTimer), userInfo: nil, repeats: true)
    }

    @objc func updateTimer(){
        print("AQUII",connectionAttempts)
    if uStatusFlag == true{
        if uStatus == "Success" {
             SwiftSpinner.sharedInstance.outerColor = UIColor.green.withAlphaComponent(0.5)
            uStatusFlag = false
            SwiftSpinner.hide()
            self.performSegue(withIdentifier: "logInSegue", sender: nil)
            stopTimer()
           
        }
        else if connectionAttempts == 4{
            print("Username or Password did not match")
            connectionAttempts = 0
            SwiftSpinner.hide()
            stopTimer()
        }
        else if connectionAttempts >= 3{
            
            SwiftSpinner.sharedInstance.outerColor = UIColor.red.withAlphaComponent(0.5)
            SwiftSpinner.show("Wrong User or Password!", animated: false)
            connectionAttempts += 1
        }
        else{
            connectionAttempts += 1
        }
    }
    }
    func stopTimer() {
        if timer != nil {
            timer?.invalidate()
            timer = nil
        }
    }
    /*
    // MARK: - Navigation

    // In a storyboard-based application, you will often want to do a little preparation before navigation
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        // Get the new view controller using segue.destinationViewController.
        // Pass the selected object to the new view controller.
    }
    */

}
