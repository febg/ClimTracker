//
//  FirstViewController.swift
//  ClimbTracker
//
//  Created by Felipe Ballesteros on 2017-10-06.
//  Copyright © 2017 Felipe Ballesteros. All rights reserved.
//

import UIKit

class FirstViewController: UIViewController {

    var time = 0
    var id = 0
    var forthPass = false
    var level = 100
    var timer2:                  Timer?
    var timer:                  Timer?
    
    @IBOutlet weak var centerH: UIImageView!
    @IBOutlet weak var upperH: UIImageView!
    @IBOutlet weak var upperRightH: UIImageView!
    @IBOutlet weak var upperLeft: UIImageView!
    @IBOutlet weak var lowerLeftH: UIImageView!
    @IBOutlet weak var lowerRightH: UIImageView!
    @IBOutlet weak var Hexagon: UIImageView!
    @IBOutlet weak var lowerH: UIImageView!
    @IBOutlet weak var whiteH: UIImageView!
    @IBOutlet weak var blackH: UIImageView!
    @IBOutlet weak var greenH: UIImageView!
    @IBOutlet weak var checkInB: UIButton!
    @IBOutlet weak var redH: UIImageView!
    @IBOutlet weak var purpleH: UIImageView!
    @IBOutlet weak var orangeH: UIImageView!
    @IBOutlet weak var yellowH: UIImageView!
    @IBOutlet weak var blueH: UIImageView!
    @IBAction func checkIn(_ sender: Any) {
        time = level
        if (level != 100 ){
            runTimer()
            checkInB.alpha = 0.15
            centerH.alpha = 0.65
            runTimer2()
            Getrequest()
            level = 100
            id = 1
        }
    }
   
    override func viewDidLoad() {
        super.viewDidLoad()
         postRequest("http://localhost:8080/testPost/")
        initializeObjects()
        
    }
    func initializeObjects(){
        let purpleGestureRecognizer = UITapGestureRecognizer(target: self, action: #selector(self.purpleHanlder))
        let orangeGestureRecognizer = UITapGestureRecognizer(target: self, action: #selector(self.orangeHanlder))
        let redGestureRecognizer = UITapGestureRecognizer(target: self, action: #selector(self.redHanlder))
        let blackGestureRecognizer = UITapGestureRecognizer(target: self, action: #selector(self.blackHanlder))
        let greenGestureRecognizer = UITapGestureRecognizer(target: self, action: #selector(self.greenHanlder))
        let blueGestureRecognizer = UITapGestureRecognizer(target: self, action: #selector(self.blueHanlder))
        let whiteGestureRecognizer = UITapGestureRecognizer(target: self, action: #selector(self.whiteHanlder))
        let yellowGestureRecognizer = UITapGestureRecognizer(target: self, action: #selector(self.yellowHanlder))
        orangeH.addGestureRecognizer(orangeGestureRecognizer)
        whiteH.addGestureRecognizer(whiteGestureRecognizer)
        blackH.addGestureRecognizer(blackGestureRecognizer)
        blueH.addGestureRecognizer(blueGestureRecognizer)
        greenH.addGestureRecognizer(greenGestureRecognizer)
        yellowH.addGestureRecognizer(yellowGestureRecognizer)
        redH.addGestureRecognizer(redGestureRecognizer)
        purpleH.addGestureRecognizer(purpleGestureRecognizer)
    }
    
    func runTimer() {
        timer = Timer.scheduledTimer(timeInterval: 0.055, target:self, selector: #selector(updateTimer), userInfo: nil, repeats: true)
    }
    func runTimer2() {
        timer2 = Timer.scheduledTimer(timeInterval: 0.06, target:self, selector: #selector(updateTimer2), userInfo: nil, repeats: true)
    }
    func stopTimer() {
        if timer != nil {
            timer?.invalidate()
            timer = nil
        }
    }
    func stopTimer2() {
        if timer2 != nil {
            
            timer2?.invalidate()
            timer2 = nil
        }
    }
    @objc func updateTimer2(){
        
         if (id == 1
            ){
        centerH.alpha = 0.15
        stopTimer2()
            id = 0
        }
    }
    @objc func updateTimer(){
        switch time {
        case 1: self.upperH.alpha = 0.0
            self.lowerH.alpha = 0.0
            self.upperRightH.alpha = 0.0
            self.lowerH.alpha = 0.0
            self.lowerRightH.alpha = 0.0
            self.lowerLeftH.alpha = 0.0
            self.upperLeft.alpha = 0.0
            stopTimer()
        
            time = time - 1
        case 2:self.upperH.alpha = 1.0
                self.lowerH.alpha = 0.0
                self.upperRightH.alpha = 0.0
                self.lowerH.alpha = 0.0
                self.lowerRightH.alpha = 0.0
                self.lowerLeftH.alpha = 0.0
                self.upperLeft.alpha = 0.0
        
            time = time - 1
        case 3: self.upperH.alpha = 1.0
                 self.lowerH.alpha = 0.0
                 self.upperRightH.alpha = 1.0
                 self.lowerH.alpha = 0.0
                 self.lowerRightH.alpha = 0.0
                 self.lowerLeftH.alpha = 0.0
                 self.upperLeft.alpha = 0.0
                time = time - 1
        case 4:self.upperH.alpha = 1.0
                self.lowerH.alpha = 0.0
                self.upperRightH.alpha = 1.0
                self.lowerH.alpha = 0.0
                self.lowerRightH.alpha = 1.0
                self.lowerLeftH.alpha = 0.0
                self.upperLeft.alpha = 0.0
                time = time - 1
        case 5:self.upperH.alpha = 1.0
                self.lowerH.alpha = 1.0
                self.upperRightH.alpha = 1.0
                self.lowerH.alpha = 1.0
                self.lowerRightH.alpha = 1.0
                self.lowerLeftH.alpha = 0.0
                self.upperLeft.alpha = 0.0
                time = time - 1
        case 6:self.upperH.alpha = 1.0
                self.lowerH.alpha = 1.0
                self.upperRightH.alpha = 1.0
                self.lowerH.alpha = 1.0
                self.lowerRightH.alpha = 1.0
                self.lowerLeftH.alpha = 1.0
                self.upperLeft.alpha = 0.0
        
            time = time - 1
        default:
            self.upperH.alpha = 0.0
            self.lowerH.alpha = 0.0
            self.upperRightH.alpha = 0.0
            self.lowerH.alpha = 0.0
            self.lowerRightH.alpha = 0.0
            self.lowerLeftH.alpha = 0.0
            self.upperLeft.alpha = 0.0
           
            time = time - 1
        }
    }
    
    func Getrequest(){
        print("Entered request func")
        let parameter = String(level)
        let url = URL(string: "http://standard-lb-1065564336.us-east-2.elb.amazonaws.com/addBlock/?level=\(parameter)&userId=D1E85310-FDCB-49E1-A13E-7E7AC790A646")
        

        let task = URLSession.shared.dataTask(with: url!) { (data, response, error) in
            
            if let data = data {
                print(type(of: data))
                print(response)
            } else if let error = error {
                print(error.localizedDescription)
            }
            
        }
        
       
        
        task.resume()
        
        // Infinitely run the main loop to wait for our request.
        // Only necessary if you are testing in the command line.
     
    
    }

    
    @objc func orangeHanlder(sender: UITapGestureRecognizer) {
        self.upperH.image = orangeH.image!
        self.lowerH.image = orangeH.image!
        self.upperRightH.image = orangeH.image!
        self.lowerH.image = orangeH.image!
        self.lowerRightH.image = orangeH.image!
        self.lowerLeftH.image = orangeH.image!
        self.upperLeft.image = orangeH.image!
    }
    @objc func purpleHanlder(sender: UITapGestureRecognizer) {
        self.upperH.image = purpleH.image!
        self.lowerH.image = purpleH.image!
        self.upperRightH.image = purpleH.image!
        self.lowerH.image = purpleH.image!
        self.lowerRightH.image = purpleH.image!
        self.lowerLeftH.image = purpleH.image!
        self.upperLeft.image = purpleH.image!
        
    }
    @objc func redHanlder(sender: UITapGestureRecognizer) {
        self.upperH.image = redH.image!
        self.lowerH.image = redH.image!
        self.upperRightH.image = redH.image!
        self.lowerH.image = redH.image!
        self.lowerRightH.image = redH.image!
        self.lowerLeftH.image = redH.image!
        self.upperLeft.image = redH.image!
        
    }
    @objc func blueHanlder(sender: UITapGestureRecognizer) {
        self.upperH.image = blueH.image!
        self.lowerH.image = blueH.image!
        self.upperRightH.image = blueH.image!
        self.lowerH.image = blueH.image!
        self.lowerRightH.image = blueH.image!
        self.lowerLeftH.image = blueH.image!
        self.upperLeft.image = blueH.image!
        
    }
    @objc func greenHanlder(sender: UITapGestureRecognizer) {
        self.upperH.image = greenH.image!
        self.lowerH.image = greenH.image!
        self.upperRightH.image = greenH.image!
        self.lowerH.image = greenH.image!
        self.lowerRightH.image = greenH.image!
        self.lowerLeftH.image = greenH.image!
        self.upperLeft.image = greenH.image!
        
    }
    @objc func blackHanlder(sender: UITapGestureRecognizer) {
        self.upperH.image = blackH.image!
        self.lowerH.image = blackH.image!
        self.upperRightH.image = blackH.image!
        self.lowerH.image = blackH.image!
        self.lowerRightH.image = blackH.image!
        self.lowerLeftH.image = blackH.image!
        self.upperLeft.image = blackH.image!
        
    }
    
    @objc func whiteHanlder(sender: UITapGestureRecognizer) {
        self.upperH.image = whiteH.image!
        self.lowerH.image = whiteH.image!
        self.upperRightH.image = whiteH.image!
        self.lowerH.image = whiteH.image!
        self.lowerRightH.image = whiteH.image!
        self.lowerLeftH.image = whiteH.image!
        self.upperLeft.image = whiteH.image!
        
    }
    
    @objc func yellowHanlder(sender: UITapGestureRecognizer) {
        self.upperH.image = yellowH.image!
        self.lowerH.image = yellowH.image!
        self.upperRightH.image = yellowH.image!
        self.lowerH.image = yellowH.image!
        self.lowerRightH.image = yellowH.image!
        self.lowerLeftH.image = yellowH.image!
        self.upperLeft.image = yellowH.image!
        
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }

    override func touchesMoved(_ touches: Set<UITouch>, with event: UIEvent?) {
        let touch: UITouch = touches.first as! UITouch
        
        var xPosition = Float(touch.location(in: self.view).x)
        var yPosition = Float(touch.location(in: self.view).y)
        
        if level != 100 {
                checkInB.alpha = 0.55
        }else{
            checkInB.alpha = 0.15
        }

        if (Float(xPosition) >= Float(upperH.frame.minX) && Float(xPosition) <= Float(upperH.frame.maxX) && Float(yPosition) >= Float(upperH.frame.minY) && Float(yPosition) <= Float(upperH.frame.maxY)){
            self.upperH.alpha = 1.0
            self.lowerH.alpha = 0.0
            self.upperRightH.alpha = 0.0
            self.lowerH.alpha = 0.0
            self.lowerRightH.alpha = 0.0
            self.lowerLeftH.alpha = 0.0
            self.upperLeft.alpha = 0.0
            level = 1
        }
       
        
        if (Float(xPosition) >= Float(upperRightH.frame.minX) && Float(xPosition) <= Float(upperRightH.frame.maxX) && Float(yPosition) >= Float(upperRightH.frame.minY) && Float(yPosition) <= Float(upperRightH.frame.maxY)){
            
            self.upperH.alpha = 1.0
            self.lowerH.alpha = 0.0
            self.upperRightH.alpha = 1.0
            self.lowerH.alpha = 0.0
            self.lowerRightH.alpha = 0.0
            self.lowerLeftH.alpha = 0.0
            self.upperLeft.alpha = 0.0
            level = 2
        }
        if ( Float(xPosition) <= Float(lowerRightH.frame.maxX) && Float(yPosition) >= Float(lowerRightH.frame.minY) && Float(xPosition) >= Float(lowerRightH.frame.minX) && Float(yPosition) <= Float(lowerRightH.frame.maxY)){
            self.upperH.alpha = 1.0
            self.lowerH.alpha = 0.0
            self.upperRightH.alpha = 1.0
            self.lowerRightH.alpha = 1.0
            self.lowerLeftH.alpha = 0.0
            self.upperLeft.alpha = 0.0
            level = 3
        }
        if ( Float(xPosition) <= Float(lowerH.frame.maxX) && Float(yPosition) >= Float(lowerH.frame.minY) && Float(xPosition) >= Float(lowerH.frame.minX) && Float(yPosition) <= Float(lowerH.frame.maxY)){
            self.upperH.alpha = 1.0
            self.lowerH.alpha = 1.0
            self.upperRightH.alpha = 1.0
            self.lowerRightH.alpha = 1.0
            self.lowerLeftH.alpha = 0.0
            self.upperLeft.alpha = 0.0
            level = 4
        }
        if ( Float(xPosition) <= Float(lowerLeftH.frame.maxX) && Float(yPosition) >= Float(lowerLeftH.frame.minY) && Float(xPosition) >= Float(lowerLeftH.frame.minX) && Float(yPosition) <= Float(lowerLeftH.frame.maxY)){
            self.upperH.alpha = 1.0
            self.lowerH.alpha = 1.0
            self.upperRightH.alpha = 1.0
            self.lowerH.alpha = 1.0
            self.lowerRightH.alpha = 1.0
            self.lowerLeftH.alpha = 1.0
            self.upperLeft.alpha = 0.0
            level = 5
        }
        if ( Float(xPosition) <= Float(upperLeft.frame.maxX) && Float(yPosition) >= Float(upperLeft.frame.minY) && Float(xPosition) >= Float(upperLeft.frame.minX) && Float(yPosition) <= Float(upperLeft.frame.maxY)){
            self.upperH.alpha = 1.0
            self.lowerH.alpha = 1.0
            self.upperRightH.alpha = 1.0
            self.lowerH.alpha = 1.0
            self.lowerRightH.alpha = 1.0
            self.lowerLeftH.alpha = 1.0
            self.upperLeft.alpha = 1.0
            level = 6
        }
        
    }
    
    override func touchesBegan(_ touches: Set<UITouch>, with event: UIEvent?) {
        let touch: UITouch = touches.first as! UITouch
        var xPosition = Float(touch.location(in: self.view).x)
        var yPosition = Float(touch.location(in: self.view).y)
        
        
        if (Float(xPosition) >= Float(upperH.frame.minX) && Float(xPosition) <= Float(upperH.frame.maxX) && Float(yPosition) >= Float(upperH.frame.minY) && Float(yPosition) <= Float(upperH.frame.maxY)){
            self.upperH.alpha = 1.0
            self.lowerH.alpha = 0.0
            self.upperRightH.alpha = 0.0
            self.lowerH.alpha = 0.0
            self.lowerRightH.alpha = 0.0
            self.lowerLeftH.alpha = 0.0
            self.upperLeft.alpha = 0.0
            level = 1
            print("Level2", level)
        }
        
        
        if (Float(xPosition) >= Float(upperRightH.frame.minX) && Float(xPosition) <= Float(upperRightH.frame.maxX) && Float(yPosition) >= Float(upperRightH.frame.minY) && Float(yPosition) <= Float(upperRightH.frame.maxY)){
            
            self.upperH.alpha = 1.0
            self.lowerH.alpha = 0.0
            self.upperRightH.alpha = 1.0
            self.lowerH.alpha = 0.0
            self.lowerRightH.alpha = 0.0
            self.lowerLeftH.alpha = 0.0
            self.upperLeft.alpha = 0.0
            level = 2
        }
        if ( Float(xPosition) <= Float(lowerRightH.frame.maxX) && Float(yPosition) >= Float(lowerRightH.frame.minY) && Float(xPosition) >= Float(lowerRightH.frame.minX) && Float(yPosition) <= Float(lowerRightH.frame.maxY)){
            self.upperH.alpha = 1.0
            self.lowerH.alpha = 0.0
            self.upperRightH.alpha = 1.0
            self.lowerRightH.alpha = 1.0
            self.lowerLeftH.alpha = 0.0
            self.upperLeft.alpha = 0.0
            level = 3
        }
        if ( Float(xPosition) <= Float(lowerH.frame.maxX) && Float(yPosition) >= Float(lowerH.frame.minY) && Float(xPosition) >= Float(lowerH.frame.minX) && Float(yPosition) <= Float(lowerH.frame.maxY)){
            self.upperH.alpha = 1.0
            self.lowerH.alpha = 1.0
            self.upperRightH.alpha = 1.0
            self.lowerRightH.alpha = 1.0
            self.lowerLeftH.alpha = 0.0
            self.upperLeft.alpha = 0.0
            level = 4
        }
        if ( Float(xPosition) <= Float(lowerLeftH.frame.maxX) && Float(yPosition) >= Float(lowerLeftH.frame.minY) && Float(xPosition) >= Float(lowerLeftH.frame.minX) && Float(yPosition) <= Float(lowerLeftH.frame.maxY)){
            self.upperH.alpha = 1.0
            self.lowerH.alpha = 1.0
            self.upperRightH.alpha = 1.0
            self.lowerH.alpha = 1.0
            self.lowerRightH.alpha = 1.0
            self.lowerLeftH.alpha = 1.0
            self.upperLeft.alpha = 0.0
            level = 5
        }
        if ( Float(xPosition) <= Float(upperLeft.frame.maxX) && Float(yPosition) >= Float(upperLeft.frame.minY) && Float(xPosition) >= Float(upperLeft.frame.minX) && Float(yPosition) <= Float(upperLeft.frame.maxY)){
            self.upperH.alpha = 1.0
            self.lowerH.alpha = 1.0
            self.upperRightH.alpha = 1.0
            self.lowerH.alpha = 1.0
            self.lowerRightH.alpha = 1.0
            self.lowerLeftH.alpha = 1.0
            self.upperLeft.alpha = 1.0
            level = 6
        }
        if level != 100 {
            checkInB.alpha = 0.55
        }else{
            checkInB.alpha = 0.15
        }
    }
    func postRequest(_ url:String)
    {
        let url:NSURL = NSURL(string: url)!
        let session = URLSession.shared
        
        let request = NSMutableURLRequest(url: url as URL)
        request.httpMethod = "POST"
        
        let params = "email=felipeb85@gmail.com&pass=test2&name=felipe5"
        //let pass = "pass=test"
        request.httpBody = params.data(using: String.Encoding.utf8)
        //request.httpBody = pass.data(using: String.Encoding.utf8)
        
        let task = session.dataTask(with: request as URLRequest) {
            (
            data, response, error) in
            if let data = data,
                let urlContent = NSString(data: data, encoding: String.Encoding.ascii.rawValue) {
                print(urlContent)
            } else {
                print("Error: \(error)")
            }
            
        }
        
        task.resume()
        
    }

}

