//
//  SignUpViewController.swift
//  ClimbTracker
//
//  Created by Felipe Ballesteros on 2017-10-10.
//  Copyright © 2017 Felipe Ballesteros. All rights reserved.
//

import UIKit

class SignUpViewController: UIViewController {

    var uName = ""
    var uPassword = ""
    var uEmail = ""
    var status = ""
    @IBOutlet weak var name: UITextField!
    @IBOutlet weak var password: UITextField!
    @IBOutlet weak var email: UITextField!
    
    
    @IBAction func signUp(_ sender: Any) {
        if (email.text! != "" && password.text! != "" && name.text! != "") {
            uEmail = email.text!
            uPassword = password.text!
            uName = name.text!
            postRequest("http://standard-lb-1065564336.us-east-2.elb.amazonaws.com/register/")
            password.text = ""
        }
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
        
        let params = "email="+uEmail+"&password="+uPassword+"&name="+uName
        //let pass = "pass=test"
        request.httpBody = params.data(using: String.Encoding.utf8)
        //request.httpBody = pass.data(using: String.Encoding.utf8)
        
        let task = session.dataTask(with: request as URLRequest) {
            (
            data, response, error) in
            if let data = data,
                let urlContent = NSString(data: data, encoding: String.Encoding.ascii.rawValue) {
               //print(type(of: urlContent))
            } else {
                print("Error: \(error)")
            }
            
        }
        
        task.resume()
        
    }
    func checkResponse(){
        
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
