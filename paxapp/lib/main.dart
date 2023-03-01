import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:paxapp/work.dart';
import 'package:pin_code_fields/pin_code_fields.dart';
import 'package:http/http.dart' as http;

void main() {
  runApp(MaterialApp(
    debugShowCheckedModeBanner: false,
    home: Home(),
  ));
}

class Home extends StatefulWidget {
  const Home({super.key});

  @override
  State<Home> createState() => _HomeState();
}

void getUser(context, String value) async {
  var url = Uri.http('213.27.32.24:8000', 'user/' + value);
  print(url);
  var response = await http.post(url);
  Navigator.push(context, MaterialPageRoute(builder: (context) => App()));
}

class _HomeState extends State<Home> {
  TextEditingController controller = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        centerTitle: true,
        title: Text("AVIAIT pax control"),
        backgroundColor: Colors.lightBlueAccent,
      ),
      body: content(context),
    );
  }

  Widget content(BuildContext context) {
    return SingleChildScrollView(
        child: Column(mainAxisAlignment: MainAxisAlignment.center, children: [
      Container(
        padding: EdgeInsets.all(20),
        child: Image.asset(
          'assets/images/air-logo.png',
          width: 250,
          height: 250,
        ),
      ),
      const SizedBox(
        width: 100.0,
        height: 100.0,
      ),
      Container(
        margin: EdgeInsets.symmetric(horizontal: 24),
        child: PinCodeTextField(
          appContext: context,
          controller: controller,
          length: 6,
          cursorHeight: 19,
          enableActiveFill: true,
          textStyle: TextStyle(fontSize: 20, fontWeight: FontWeight.normal),
          inputFormatters: [FilteringTextInputFormatter.digitsOnly],
          pinTheme: PinTheme(
            shape: PinCodeFieldShape.underline,

            //fieldWidth: 50,
            inactiveColor: Colors.grey,
            selectedColor: Colors.lightBlueAccent,
            activeFillColor: Colors.white,
            inactiveFillColor: Colors.white,
            //borderWidth: 1
          ),
          keyboardType: TextInputType.number,
          onChanged: ((value) {
            print(value.length);
            if (value.length == 6) {
              getUser(context, value);
            }
          }),
        ),
      ),
    ]));
  }
}
